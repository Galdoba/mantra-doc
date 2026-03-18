---
updated_at: 2026-03-15T09:58:33.295+10:00
tags:
  - logger
  - modular
---
# Rotate Module - Automatic Log File Rotation

## Overview

The rotate module provides automatic log file rotation for molog. When enabled, it monitors log files and rotates them based on configured triggers (size, age, records).

## Status: 🔧 In Development

The rotate module implements periodic checking workflow. Size-based rotation is implemented. Age and records-based rotation planned.

## Configuration

### Config Structure

```go
type Config struct {
    SizeMB        int           // Max file size in MB before rotation (0 = disabled)
    AgeHours      int           // Max file age in hours before rotation (0 = disabled)
    RecordsInFile int           // Max records per file (0 = disabled)
    Archive       Archive       // Archive configuration
    CheckInterval time.Duration // Rotation check interval (0 = default 10 min, min 1 min)
}
```

### Archive Structure

```go
type Archive struct {
    Directory     string // Directory to store rotated files
    FilePrefix    string // Prefix for archived files
    FileSuffix    string // Suffix for archived files
    FileExtention string // File extension for archived files
}
```

### Validation Rules

- `SizeMB`: Must be non-negative
- `AgeHours`: Must be non-negative
- `RecordsInFile`: Must be non-negative
- `CheckInterval`: Minimum 1 minute (enforced in code)

## Module Setup

### WithModuleRotate(cfg rotatemod.Config) Option

Enables automatic log file rotation with specified configuration.

```go
logFile, _ := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

logger, err := molog.New(
    molog.WithWriter("file", logFile),
    molog.WithModuleRotate(rotatemod.Config{
        SizeMB:        100,
        AgeHours:      24,
        CheckInterval: 10 * time.Minute,
    }),
)
```

**Examples:**

```go
// Rotate by size only (10 min default interval)
molog.WithModuleRotate(rotatemod.Config{
    SizeMB: 100,
})

// Rotate by age only
molog.WithModuleRotate(rotatemod.Config{
    AgeHours: 24,
})

// Rotate by records only
molog.WithModuleRotate(rotatemod.Config{
    RecordsInFile: 100000,
})

// Custom check interval (minimum 1 minute)
molog.WithModuleRotate(rotatemod.Config{
    SizeMB:        100,
    CheckInterval: 1 * time.Minute,
})

// Full config with archive
molog.WithModuleRotate(rotatemod.Config{
    SizeMB:        100,
    AgeHours:      24,
    RecordsInFile: 100000,
    CheckInterval: 10 * time.Minute,
    Archive: rotatemod.Archive{
        Directory:     "./logs/archive",
        FilePrefix:    "app",
        FileExtention: ".log",
    },
})
```

## Architecture

### Module Structure

```go
type fileInfo struct {
    file        *os.File
    createdAt   time.Time
    size        int64
    recordCount int64
}

type Module struct {
    mu           sync.Mutex
    enabled      bool
    trackedFiles map[string]fileInfo  // path -> info
    config       Config
    archive      Archive
    
    // Worker
    wakeCh       chan struct{}
    stopCh       chan struct{}
    wg           sync.WaitGroup
}
```

**Key design decisions:**
- Uses `map[string]fileInfo` to track files with metadata
- Worker-based periodic checking (not per-log-call)
- All rotation logic owned by worker
- TryRotate() wakes worker for immediate check

### Workflow

| Trigger | Action |
|---------|--------|
| Configure() | Initial sweep + start worker (immediate first run) |
| Worker timeout | Full sweep → eval triggers → rotate if needed → sleep |
| TryRotate() | Signal wakeCh → worker does full sweep on wake |

### Worker Cycle

```
Wake (on start/timeout/TryRotate)
    ↓
Full sweep: stat files, update fileInfo (size, recordCount+=1)
    ↓
Evaluate triggers (size, age, records)
    ↓
If trigger=true: rotate + reset fileInfo (createdAt=now, size=0, recordCount=0)
    ↓
Sleep CheckInterval (default 10 min)
```

### Trigger Evaluation

```go
func shouldRotate(info fileInfo, cfg Config) bool {
    // Size check
    if cfg.SizeMB > 0 && info.size > int64(cfg.SizeMB*1024*1024) {
        return true
    }
    // Age check
    if cfg.AgeHours > 0 && time.Since(info.createdAt) > time.Duration(cfg.AgeHours)*time.Hour {
        return true
    }
    // Records check
    if cfg.RecordsInFile > 0 && info.recordCount > int64(cfg.RecordsInFile) {
        return true
    }
    return false
}
```

## Runtime API

### TryRotate()

Forces an immediate rotation check. Worker will wake up and perform full sweep.

```go
logger.(*molog.loggerBase).TryRotate()
```

**Note**: Requires type assertion to access the method.

## Usage Examples

### Basic rotation by size

```go
logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
if err != nil {
    panic(err)
}

logger, err := molog.New(
    molog.WithWriter("file", logFile),
    molog.WithHandlerType(molog.HANDLER_TEXT),
    molog.WithModuleRotate(rotatemod.Config{
        SizeMB: 100,
    }),
)
if err != nil {
    panic(err)
}
```

### Async with rotation

The rotate module is thread-safe and works with async logging:

```go
logFile, _ := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

logger, _ := molog.New(
    molog.WithWriter("file", logFile),
    molog.WithHandlerType(molog.HANDLER_TEXT),
    molog.WithModuleAsync(asyncmod.Config{
        Routines: 2,
        Buffer:   1024,
    }),
    molog.WithModuleRotate(rotatemod.Config{
        SizeMB:        100,
        CheckInterval: 10 * time.Minute,
    }),
)
```

### Force rotation

```go
lb := logger.(*molog.loggerBase)
lb.TryRotate()  // Wake worker for immediate check
```

## Implementation Status

### Implemented

- [x] Config structure with CheckInterval
- [x] Validate() method
- [x] Module creation with file tracking (map-based)
- [x] Configure() method with worker startup
- [x] Worker-based periodic checking
- [x] Size-based rotation check and execution
- [x] Thread safety (mutex + worker) for async compatibility

### Planned

- [ ] Age-based rotation
- [ ] Record count rotation
- [ ] TryRotate() method
- [ ] Archive directory creation
- [ ] Compression (gzip)
- [ ] Retention policies

## Design Rationale

### Why Periodic Checking?

For small to medium applications, checking file stats on every log call is unnecessary overhead. A 10-minute default interval is appropriate because:

1. Log files don't grow to rotation threshold in seconds for most apps
2. If rotation is needed more frequently, consider a different logging solution
3. Significantly reduces syscall overhead
4. Age-based rotation doesn't need second-level precision

### Why Worker-Owns-All?

All rotation logic is centralized in the worker:

- Single source of truth for trigger evaluation
- Simpler thread safety (one goroutine does everything)
- TryRotate() just signals - no duplicate evaluation
- Easier to reason about state

See [rotationRoadmap.md](rotationRoadmap.md) for detailed implementation plan.

## Test Coverage

Currently no tests exist for the rotate module.

Run tests when implemented:

```bash
go test -v ./modules/rotatemod/...
```
