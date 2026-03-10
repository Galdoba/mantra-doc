---
updated_at: 2026-03-05T04:47:05.811+10:00
tags:
---
# Async Module - Non-Blocking Logger Implementation

## Overview

This document describes the implementation plan for the non-blocking (async) logging mode in molog. When enabled, log records are sent to a channel and processed by background worker routines instead of blocking the main thread.

## Architecture

### Core Idea

When async mode is enabled:
1. The main logging call (`log()` / `logAttrs()`) creates a `slog.Record` and sends it to a channel
2. Background worker routines read from the channel and write to the handler
3. The main thread returns immediately without blocking

When async mode is disabled:
1. The main logging call processes the record synchronously
2. The handler writes to the output immediately (blocking behavior)

## Data Structures

### Configuration Structure (asyncModule)

```go
type asyncModule struct {
    enabled      bool           // whether async mode is active
    routines     int            // number of worker routines (default: 1)
    buffer       int            // channel buffer size (default: 0 - unbuffered)
    logOverflow  bool           // log when channel buffer is full
    mu           sync.Mutex     // protects enabled state
    recordCh     chan Record    // channel for passing records to workers
    stopCh       chan struct{}  // signal to stop workers
}
```

### Placement

The `asyncModule` struct is embedded within the `modulesConfiguration`:

```go
type modulesConfiguration struct {
    base         base
    asyncModule asyncModule  // NEW - async configuration
}
```

Note: The asyncModule is at the modulesConfiguration level, not inside base, to keep it separate from the core logger configuration.

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
func (l *loggerBase) startWorkers() {
    // Create channels
    l.cfg.asyncModule.recordCh = make(chan Record, l.cfg.asyncModule.buffer)
    l.cfg.asyncModule.stopCh = make(chan struct{})
    
    // Start worker routines
    for i := 0; i < l.cfg.asyncModule.routines; i++ {
        go l.worker()
    }
}

func (l *loggerBase) worker() {
    for {
        select {
        case record := <-l.cfg.asyncModule.recordCh:
            // Process the record
            l.slog.Handler().Handle(context.Background(), record)
        case <-l.cfg.asyncModule.stopCh:
            // Drain remaining records then exit
            for {
                select {
                case record := <-l.cfg.asyncModule.recordCh:
                    l.slog.Handler().Handle(context.Background(), record)
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
func (l *loggerBase) ToggleAsync(enabled bool) {
    l.cfg.asyncModule.mu.Lock()
    defer l.cfg.asyncModule.mu.Unlock()
    
    // Already in desired state
    if enabled == l.cfg.asyncModule.enabled {
        return
    }
    
    if enabled {
        // Enable: start workers
        l.startWorkers()
    } else {
        // Disable: stop workers
        l.stopWorkers()
    }
    
    l.cfg.asyncModule.enabled = enabled
}
```

## AsyncReconfigure Implementation

```go
func (l *loggerBase) AsyncReconfigure(params ...int) {
    l.cfg.asyncModule.mu.Lock()
    defer l.cfg.asyncModule.mu.Unlock()
    
    // Parse parameters
    newRoutines := l.cfg.asyncModule.routines
    newBuffer := l.cfg.asyncModule.buffer
    
    if len(params) > 0 && params[0] > 0 {
        newRoutines = params[0]
    }
    if len(params) > 1 && params[1] > 0 {
        newBuffer = params[1]
    }
    
    // If async is currently enabled, restart with new settings
    if l.cfg.asyncModule.enabled {
        l.stopWorkers()
        l.cfg.asyncModule.routines = newRoutines
        l.cfg.asyncModule.buffer = newBuffer
        l.startWorkers()
    } else {
        // Just update configuration for future enable
        l.cfg.asyncModule.routines = newRoutines
        l.cfg.asyncModule.buffer = newBuffer
    }
}
```

## Error Handling

### Channel Overflow

When the channel buffer is full:

```go
select {
case l.cfg.asyncModule.recordCh <- r:
    // Sent successfully
default:
    // Buffer full
    if l.cfg.asyncModule.logOverflow {
        // Log overflow notification (via original handler, not async)
        // This prevents infinite recursion
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

- [ ] Add `asyncModule` struct to `config.go`
- [ ] Add `asyncModule` field to `modulesConfiguration` struct in `config.go`
- [ ] Add `WithAsync` option function in `options.go`
- [ ] Implement `startWorkers()` method in `baseLogger.go`
- [ ] Implement `stopWorkers()` method in `baseLogger.go`
- [ ] Implement `worker()` method in `baseLogger.go`
- [ ] Modify `log()` to support async mode in `baseLogger.go`
- [ ] Modify `logAttrs()` to support async mode in `baseLogger.go`
- [ ] Implement `ToggleAsync()` method in `baseLogger.go`
- [ ] Implement `AsyncReconfigure()` method in `baseLogger.go`
- [ ] Add tests for async functionality
- [ ] Update documentation

## Notes

- The async mode is designed for high-throughput scenarios where logging should not block the main thread
- With multiple workers, log records may be written out of order
- The buffer size of 0 (unbuffered) provides the lowest latency but workers must keep up with the rate
- Larger buffer sizes can handle burst traffic but may increase memory usage
- When disabled, all pending records in the channel are processed before switching to blocking mode
- The asyncModule is placed at the modulesConfiguration level to allow for future module additions
