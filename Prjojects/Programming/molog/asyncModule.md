---
updated_at: 2026-03-15T09:58:02.028+10:00
tags:
  - logger
  - modular
---
# Async Module - Non-Blocking Logger Implementation

## Overview

The async module provides non-blocking logging for molog. When enabled, log records are sent to a channel and processed by background worker routines instead of blocking the main thread.

## Status: ✅ Stable

The async module is fully implemented with comprehensive test coverage.

## Configuration

### Config Structure

```go
type Config struct {
    Routines    int  // Number of worker routines (1-10)
    Buffer      int  // Channel buffer size (0 = unbuffered)
    LogOverflow bool // Log warning when buffer is full
}
```

### Validation Rules

- `Routines`: Must be between 1 and 10
- `Buffer`: Must be non-negative

## Module Setup

### WithModuleAsync(cfg asyncmod.Config) Option

Enables async logging with specified configuration.

```go
logger, err := molog.New(
    molog.WithModuleAsync(asyncmod.Config{
        Routines:    2,
        Buffer:      1024,
        LogOverflow: true,
    }),
)
```

**Examples:**

```go
// Minimal config (1 routine, unbuffered)
molog.WithModuleAsync(asyncmod.Config{Routines: 1, Buffer: 0})

// Multiple workers with buffer
molog.WithModuleAsync(asyncmod.Config{Routines: 4, Buffer: 1024})

// With overflow logging
molog.WithModuleAsync(asyncmod.Config{
    Routines:    2,
    Buffer:      100,
    LogOverflow: true,
})
```

## Architecture

### Core Idea

When async mode is enabled:
1. The main logging call (`log()` / `logAttrs()`) creates a `slog.Record` and sends it to a channel
2. Background worker routines read from the channel and write to the handler
3. The main thread returns immediately without blocking

When async mode is disabled:
1. The async module returns `false` to indicate async should not be used
2. The base logger processes the record synchronously

### Module Structure

```go
type Module struct {
    enabled      bool             // whether async mode is active
    routines     int              // number of worker routines
    buffer       int              // channel buffer size
    logOverflow  bool             // log when channel buffer is full
    mu           sync.Mutex       // protects enabled state
    recordCh     chan slog.Record // channel for passing records to workers
    stopCh       chan struct{}    // signal to stop workers
    wg           sync.WaitGroup   // waits workers to finish
}
```

## Runtime API

### ToggleAsync(enabled bool)

Toggle async mode at runtime. Must type-assert to access:

```go
lb := logger.(*molog.loggerBase)
lb.ToggleAsync(true)   // Enable async
lb.ToggleAsync(false) // Disable async
```

**Behavior:**
- When enabling: starts worker routines if not already running
- When disabling: stops worker routines and waits for channel to drain

### AsyncReconfigure(cfg asyncmod.Config) error

Reconfigure async module at runtime:

```go
lb := logger.(*molog.loggerBase)
lb.AsyncReconfigure(asyncmod.Config{
    Routines: 4,
    Buffer:   2048,
})
```

**Note:** Changing parameters will restart workers to apply new settings.

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
logger, err := molog.New(
    molog.WithWriter("console", os.Stderr),
    molog.WithModuleAsync(asyncmod.Config{
        Routines: 1,
        Buffer:   0,
    }),
)

logger.Info("this returns immediately")
logger.Error("this also returns immediately")
```

### Multiple workers with buffer

```go
logger, _ := molog.New(
    molog.WithWriter("file", logFile),
    molog.WithHandlerType(molog.HANDLER_TEXT),
    molog.WithModuleAsync(asyncmod.Config{
        Routines:    4,
        Buffer:      1024,
        LogOverflow: true,
    }),
)
```

### Toggle at runtime

```go
logger, _ := molog.New(molog.WithLevel(molog.LevelInfo))

lb := logger.(*molog.loggerBase)

// Later, enable async mode
lb.ToggleAsync(true)

// Or disable
lb.ToggleAsync(false)
```

### Reconfigure at runtime

```go
logger, _ := molog.New(
    molog.WithWriter("file", logFile),
    molog.WithModuleAsync(asyncmod.Config{
        Routines: 2,
        Buffer:   512,
    }),
)

// Later, increase capacity
lb := logger.(*molog.loggerBase)
lb.AsyncReconfigure(asyncmod.Config{
    Routines: 8,
    Buffer:   4096,
})
```

## Implementation Notes

- Async mode is designed for high-throughput scenarios where logging should not block the main thread
- With multiple workers, log records may be written out of order
- Buffer size of 0 (unbuffered) provides lowest latency but workers must keep up with rate
- Larger buffer sizes handle burst traffic but increase memory usage
- When disabled, all pending records in the channel are processed before switching to blocking mode

## Test Coverage

The module includes tests for:

- Config validation
- Module creation and configuration
- Start/stop toggle behavior
- Record processing
- Level filtering
- Overflow handling
- Multiple concurrent routines
- Nil context handling
- Drain on stop

Run tests:

```bash
go test -v ./modules/asyncmod/...
```
