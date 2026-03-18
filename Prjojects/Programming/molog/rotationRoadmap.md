---
updated_at: 2026-03-15T09:58:40.178+10:00
tags:
  - logger
  - modular
---
# Rotation Module - Implementation Roadmap

This document outlines the detailed implementation plan for the rotate module's periodic checking workflow.

## Current State

The current implementation calls `Rotate()` after every log write, checking file size each time. This is inefficient for high-throughput logging.

## Target Design

See [rotateModule.md](rotateModule.md) for the target architecture.

## Implementation Phases

### Phase 1: Core Infrastructure

**Goal**: Add worker-based periodic checking without changing external API

#### 1.1 Update Config
```go
type Config struct {
    SizeMB        int
    AgeHours      int
    RecordsInFile int
    Archive       Archive
    CheckInterval time.Duration // NEW: 0 = default 10 min
}
```

#### 1.2 Update fileInfo
```go
type fileInfo struct {
    file        *os.File
    createdAt   time.Time  // NEW
    size        int64
    recordCount int64      // NEW (for future)
}
```

#### 1.3 Update Module
```go
type Module struct {
    mu           sync.Mutex
    enabled      bool
    trackedFiles map[string]fileInfo  // Changed: value not pointer
    config       Config
    archive      Archive
    
    // Worker channels
    wakeCh       chan struct{}
    stopCh       chan struct{}
    wg           sync.WaitGroup
}
```

#### 1.4 Implement NewWorker
```go
func (m *Module) startWorker() {
    m.wg.Add(1)
    go m.worker()
}

func (m *Module) worker() {
    defer m.wg.Done()
    
    interval := m.config.CheckInterval
    if interval == 0 {
        interval = 10 * time.Minute
    }
    if interval < time.Minute {
        interval = time.Minute
    }
    
    for {
        select {
        case <-m.wakeCh:
            // Immediate check (from TryRotate)
            m.doCheck()
        case <-m.stopCh:
            return
        case <-time.After(interval):
            // Periodic check
            m.doCheck()
        }
    }
}

func (m *Module) doCheck() {
    m.mu.Lock()
    defer m.mu.Unlock()
    
    // 1. Full sweep: stat files, update fileInfo
    // 2. Evaluate triggers
    // 3. Rotate if needed
    // 4. Reset fileInfo after rotation
}
```

#### 1.5 Update Configure()
```go
func (m *Module) Configure(cfg Config) error {
    m.mu.Lock()
    defer m.mu.Unlock()
    
    if err := cfg.Validate(); err != nil {
        return fmt.Errorf("invalid %s configuration: %w", ModuleName, err)
    }
    m.config = cfg
    
    // Initial sweep
    m.collectFileInfo()
    
    // Start worker
    m.startWorker()
    
    // Trigger immediate first check
    m.wakeCh <- struct{}{}
    
    return nil
}
```

#### 1.6 Implement collectFileInfo()
```go
func (m *Module) collectFileInfo() {
    for path, info := range m.trackedFiles {
        f, err := os.OpenFile(path, os.O_RDONLY, 0)
        if err != nil {
            continue  // Skip inaccessible files
        }
        stat, _ := f.Stat()
        info.size = stat.Size()
        info.createdAt = stat.ModTime()  // Or CreationTime if available
        m.trackedFiles[path] = info
    }
}
```

### Phase 2: Trigger Evaluation

**Goal**: Implement size, age, and records checks

#### 2.1 Implement shouldRotate()
```go
func (m *Module) shouldRotate(path string) bool {
    info := m.trackedFiles[path]
    cfg := m.config
    
    if cfg.SizeMB > 0 && info.size > int64(cfg.SizeMB*1024*1024) {
        return true
    }
    if cfg.AgeHours > 0 && time.Since(info.createdAt) > time.Duration(cfg.AgeHours)*time.Hour {
        return true
    }
    if cfg.RecordsInFile > 0 && info.recordCount > int64(cfg.RecordsInFile) {
        return true
    }
    return false
}
```

#### 2.2 Update doCheck()
```go
func (m *Module) doCheck() {
    // Refresh file info
    m.collectFileInfo()
    
    // Check each file
    for path, info := range m.trackedFiles {
        if m.shouldRotate(path) {
            m.rotate(path)
            // Reset info after rotation
            m.trackedFiles[path] = fileInfo{
                file:      info.file,
                createdAt: time.Now(),
                size:      0,
                recordCount: 0,
            }
        }
    }
}
```

### Phase 3: External API

**Goal**: Add TryRotate() method

#### 3.1 Add TryRotate() to Module
```go
func (m *Module) TryRotate() {
    select {
    case m.wakeCh <- struct{}{}:
    default:
        // Channel full, check will happen on next worker wake anyway
    }
}
```

#### 3.2 Add wrapper in base package
```go
func (l *loggerBase) TryRotate() {
    if l.cfg.rotate == nil {
        return
    }
    l.cfg.rotate.TryRotate()
}
```

### Phase 4: Cleanup & Polish

**Goal**: Complete implementation

#### 4.1 Remove old Rotate() call from base.go
- The periodic worker now handles rotation
- Remove `l.cfg.rotate.Rotate()` from `log()` method

#### 4.2 Add record counting
```go
func (m *Module) doCheck() {
    // Increment record count (approximate)
    for path := range m.trackedFiles {
        info := m.trackedFiles[path]
        info.recordCount++
        m.trackedFiles[path] = info
    }
    
    // Then check triggers...
}
```

#### 4.3 Handle worker shutdown
```go
func (m *Module) Close() error {
    close(m.stopCh)
    m.wg.Wait()
    return nil
}
```

## Implementation Checklist

- [ ] Phase 1: Core Infrastructure
  - [ ] Update Config with CheckInterval
  - [ ] Update fileInfo structure
  - [ ] Update Module structure with worker channels
  - [ ] Implement worker() function
  - [ ] Update Configure() to start worker
  - [ ] Implement collectFileInfo()
  
- [ ] Phase 2: Trigger Evaluation
  - [ ] Implement shouldRotate()
  - [ ] Update doCheck() with full logic
  
- [ ] Phase 3: External API
  - [ ] Add TryRotate() to Module
  - [ ] Add TryRotate() wrapper in base
  
- [ ] Phase 4: Cleanup
  - [ ] Remove old Rotate() call from base.go
  - [ ] Add record counting
  - [ ] Add Close() method
  - [ ] Update tests

## Notes

1. **Records counting**: Approximate is acceptable. Only check at period end, not on every log. If trigger is 100000 and we rotate at 100500, that's fine for this module's purpose.

2. **Age tracking**: Use file ModTime initially. After rotation, reset to current time.

3. **TryRotate vs immediate**: TryRotate signals worker, worker does the work. This avoids duplicate evaluation and keeps all rotation logic in one place.

4. **Backward compatibility**: Keep old API working until new one is fully implemented, then remove old code.
