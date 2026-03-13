---
updated_at: 2026-03-13T09:59:10.204+10:00
tags:
  - logger
  - modular
---
# Async Module - Non-Blocking Logger Implementation

## Overview

The async module provides non-blocking logging for molog. When enabled, log records are sent to a channel and processed by background worker routines instead of blocking the main thread.

## Status: ✅ Prototype

The async module is at prototype level with comprehensive test coverage.

## Architecture

### Core Idea

When async mode is enabled:
1. The main logging call (`log()` / `logAttrs()`) creates a `slog.Record` and sends it to a channel
2. Background worker routines read from the channel and write to the handler
3. The main thread returns immediately without blocking

When async mode is disabled:
1. The main logging call returns `false` to indicate async should not be used
2. The base logger processes the record synchronously

## Data Structures

### Module Structure (asyncm.Module)

```go
type Module struct {
    enabled      bool             // whether async mode is active
    routines     int              // number of worker routines (default: 1)
    buffer      int              // channel buffer size (default: 0 - unbuffered)
    logOverflow bool             // log when channel buffer is full
    mu          sync.Mutex       // protects enabled state
    recordCh    chan slog.Record // channel for passing records to workers
    stopCh      chan struct{}    // signal to stop workers
    wg          sync.WaitGroup   // waits workers to finish
}
```

### Placement

The `asyncm.Module` is a pointer in `modulesConfiguration`:

```go
type modulesConfiguration struct {
    base  base
    async *asyncm.Module
}
```

The module is located in `internal/modules/asyncm/async.go`.

## Options API

### WithAsync

```go
// WithAsync enables non-blocking (async) mode with optional parameters.
// This option starts worker routines that process log records in the background,
// allowing the main thread to return immediately without waiting for I/O.
//
// Parameters (all optional):
//   - First: number of worker routines (default: 1, minimum: 1)
//   - Second: channel buffer size (default: 0 - unbuffered)
//   - Third+: ignored
//
// Examples:
//   - WithAsync()                    - 1 routine, unbuffered
//   - WithAsync(4)                  - 4 routines, unbuffered
//   - WithAsync(4, 1024)            - 4 routines, buffered with 1024 capacity
func WithAsync(routinesAndBuffer ...int) Option
```

**Default values:**
- `routines`: 1 (minimum 1 when enabled)
- `buffer`: 0 (unbuffered channel)

## Runtime API

### ToggleAsync

```go
// ToggleAsync enables or disables async mode at runtime.
// When enabling: starts worker routines if not already running
// When disabling: stops worker routines and waits for channel to drain
//
// Parameters:
//   - enabled: true to enable async mode, false to disable
//
// Note: When disabling, all pending records in the channel 
// will be processed before switching to blocking mode.
func (l *loggerBase) ToggleAsync(enabled bool)
```

### AsyncReconfigure

```go
// AsyncReconfigure allows changing async module parameters at runtime.
// Can modify:
//   - Number of worker routines
//   - Channel buffer size
//
// Parameters:
//   - First: new number of routines (0 or negative = keep current)
//   - Second: new buffer size (0 or negative = keep current)
//   - Third+: ignored
//
// Note: Changing parameters will restart workers to apply new settings.
// Any pending records in the channel may be lost.
//
// Examples:
//   - AsyncReconfigure()                - no changes
//   - AsyncReconfigure(8)              - change to 8 routines
//   - AsyncReconfigure(8, 2048)        - 8 routines, buffer 2048
//   - AsyncReconfigure(0, 4096)         - keep routines, change buffer to 4096
func (l *loggerBase) AsyncReconfigure(params ...int)
```

## Behavior

### When enabled (async mode)

```
User call (Debug/Info/Warn/Error)
         ↓
    log() / logAttrs()
         ↓
    Create slog.Record
         ↓
    select {
    case recordCh <- record:
        // Success: non-blocking send
    default:
        // Buffer full: drop or log overflow (depending on logOverflow setting)
    }
         ↓
    Returns immediately to caller
```

### When disabled (blocking mode)

```
User call (Debug/Info/Warn/Error)
         ↓
    log() / logAttrs()
         ↓
    Create slog.Record
         ↓
    Handler().Handle(ctx, record)
         ↓
    Returns after write completes
```

## Worker Routine Behavior

```go
func (m *Module) startWorkers(l *slog.Logger) {
    m.recordCh = make(chan slog.Record, m.buffer)
    m.stopCh = make(chan struct{})

    for i := 0; i < m.routines; i++ {
        m.wg.Add(1)
        go m.worker(l)
    }
}

func (m *Module) worker(l *slog.Logger) {
    defer m.wg.Done()

    for {
        select {
        case record := <-m.recordCh:
            l.Handler().Handle(context.Background(), record)
        case <-m.stopCh:
            // Drain remaining records then exit
            for {
                select {
                case record := <-m.recordCh:
                    l.Handler().Handle(context.Background(), record)
                default:
                    return
                }
            }
        }
    }
}
```

## ToggleAsync Implementation

```go
func (m *Module) ToggleAsync(l *slog.Logger, enabled bool) {
    m.mu.Lock()
    defer m.mu.Unlock()

    if enabled == m.enabled {
        return
    }
    m.enabled = enabled

    if m.enabled {
        m.startWorkers(l)
    } else {
        m.stopWorkers()
    }
}
```

## AsyncReconfigure Implementation

```go
func (l *loggerBase) AsyncReconfigure(params ...int) {
    if l.cfg.async == nil {
        return
    }
    initialState := l.cfg.async.IsEnabled()
    l.cfg.async.ToggleAsync(l.slog, false)
    for i, value := range params {
        switch i {
        case 0:
            l.cfg.async.SetRoutines(value)
        case 1:
            l.cfg.async.SetBuffer(value)
        case 2:
            l.cfg.async.SetLogOverflow(asyncm.LogOverflowSetting(value))
        }
    }
    l.cfg.async.ToggleAsync(l.slog, initialState)
}
```

## Error Handling

### Channel Overflow

When the channel buffer is full:

```go
select {
case m.recordCh <- r:
    // Sent successfully
default:
    // Buffer full
    if m.logOverflow {
        l.Warn("log overflow") // Log via original handler to prevent recursion
        return false
    }
}
```

## Usage Examples

### Basic async usage

```go
logger, _ := molog.New(
    molog.WithLevel(molog.LevelInfo),
    molog.WithAsync(),  // 1 routine, unbuffered
)

logger.Info("this returns immediately")
logger.Error("this also returns immediately")
```

### Multiple workers with buffer

```go
logger, _ := molog.New(
    molog.WithLevel(molog.LevelDebug),
    molog.WithAsync(4, 1024),  // 4 routines, buffer 1024
)
```

### Toggle at runtime

```go
logger, _ := molog.New(molog.WithLevel(molog.LevelInfo))

// Later, enable async mode
logger.ToggleAsync(true)

// Or disable
logger.ToggleAsync(false)
```

### Reconfigure at runtime

```go
logger, _ := molog.New(
    molog.WithAsync(2, 512),  // Start with 2 routines, buffer 512
)

// Later, increase capacity
logger.AsyncReconfigure(8, 4096)

// Or just change one parameter
logger.AsyncReconfigure(0, 8192)  // Keep 8 routines, increase buffer
```

## Implementation Checklist

- [x] Add `asyncModule` struct to `config.go`
- [x] Add `asyncModule` field to `modulesConfiguration` struct in `config.go`
- [x] Add `WithAsync` option function in `options.go`
- [x] Implement `startWorkers()` method
- [x] Implement `stopWorkers()` method
- [x] Implement `worker()` method
- [x] Modify `log()` to support async mode
- [x] Modify `logAttrs()` to support async mode
- [x] Implement `ToggleAsync()` method in `control.go`
- [x] Implement `AsyncReconfigure()` method in `control.go`
- [x] Add tests for async functionality
- [x] Update documentation

## Notes

- The async mode is designed for high-throughput scenarios where logging should not block the main thread
- With multiple workers, log records may be written out of order
- The buffer size of 0 (unbuffered) provides the lowest latency but workers must keep up with the rate
- Larger buffer sizes can handle burst traffic but may increase memory usage
- When disabled, all pending records in the channel are processed before switching to blocking mode
---
Беседа о композиции пакетов: [[Deepseek]]
