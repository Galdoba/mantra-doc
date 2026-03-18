---
<<<<<<< HEAD
updated_at: 2026-03-15T09:58:12.105+10:00
=======
updated_at: 2026-03-13T09:58:00.868+10:00
>>>>>>> origin/main
tags:
  - logger
  - modular
---
# Base Module

The base module is the core of molog, providing a wrapper around the standard library's `log/slog` package with extended configuration capabilities.

## Overview

The base module provides full compatibility with `*slog.Logger` API and adds flexible configuration options through the `New()` constructor.

## Project Structure

```
molog/
├── base.go              # Logger interface, implementation, constructors
├── config.go            # Types, constants, configuration structures
├── options.go           # Option functions for configuration
├── control.go           # Runtime control methods (ToggleAsync, AsyncReconfigure, TryRotate)
├── base_test.go         # Tests
└── modules/
    ├── asyncmod/        # Async logging module
    └── rotatemod/       # Rotation module
```

## Available Modules

| Module | Status | Description |
|--------|--------|-------------|
| base | ✅ Stable | Core logger functionality |
| async | ✅ Stable | Non-blocking async logging |
| rotate | 🔧 In Development | Automatic log file rotation |

## Available Functionality

### Logger Interface

The `Logger` interface fully covers the public API of `*slog.Logger`:

| Method | Description |
|--------|-------------|
| `Debug(msg string, args ...any)` | Log at Debug level |
| `DebugContext(ctx, msg, args)` | Log at Debug level with context |
| `Info(msg string, args ...any)` | Log at Info level |
| `InfoContext(ctx, msg, args)` | Log at Info level with context |
| `Warn(msg string, args ...any)` | Log at Warn level |
| `WarnContext(ctx, msg, args)` | Log at Warn level with context |
| `Error(msg string, args ...any)` | Log at Error level |
| `ErrorContext(ctx, msg, args)` | Log at Error level with context |
| `Log(ctx, level, msg, args)` | Log at arbitrary level |
| `LogAttrs(ctx, level, msg, attrs)` | Log with attributes |
| `Enabled(ctx, level)` | Check if level is enabled |
| `Handler()` | Get the handler |
| `With(args)` | Create a new logger with added attributes |
| `WithGroup(name)` | Create a new logger with attribute group |

### Available Types and Aliases

```go
type Level = slog.Level      // Log level
type Attr = slog.Attr        // Record attribute
type Record = slog.Record    // Log record

const (
    LevelDebug Level = slog.LevelDebug  // Debug level
    LevelInfo  Level = slog.LevelInfo    // Info level
    LevelWarn  Level = slog.LevelWarn    // Warning level
    LevelError Level = slog.LevelError   // Error level
)

type handlerType string

const (
    HANDLER_JSON   handlerType = "json"   // JSON output format
    HANDLER_TEXT   handlerType = "text"  // Text output format
    HANDLER_CUSTOM handlerType = "custom" // Custom handler (not implemented)
)
```

## Constructor New() Options

### Base Options (With*)

#### WithWriter(writerKey string, writer io.Writer) Option

Adds a writer for log output. The key cannot be empty.

```go
logger, _ := molog.New(
    molog.WithWriter("console", os.Stderr),
)
```

When multiple writers are specified, they are all combined via `io.MultiWriter`.

#### WithHandlerType(ht handlerType) Option

Sets the handler type (JSON or TEXT). Default is JSON.

```go
logger, _ := molog.New(
    molog.WithHandlerType(molog.HANDLER_TEXT),
)
```

#### WithAddSource(addSource bool) Option

Enables/disables adding source information (file:line). Default is false.

```go
logger, _ := molog.New(
    molog.WithAddSource(true),
)
```

#### WithLevel(level Level) Option

Sets the minimum log level. Default is LevelDebug.

```go
logger, _ := molog.New(
    molog.WithLevel(molog.LevelInfo),
)
```

#### WithReplaceAttr(key string, fn func([]string, slog.Attr) slog.Attr) Option

Adds a function for attribute transformation. Functions are stored in a map by key.

```go
logger, _ := molog.New(
    molog.WithReplaceAttr("timestamp", func(groups []string, a slog.Attr) slog.Attr {
        if a.Key == "time" {
            return slog.String("timestamp", a.Value.String())
        }
        return a
    }),
)
```

#### WithCustomHandler(h slog.Handler) Option

Allows passing a completely custom handler. **Note: Currently not implemented.**

```go
// Currently not available
```

### Module Options (WithModule*)

#### WithModuleAsync(cfg asyncmod.Config) Option

Enables async logging with specified configuration.

```go
logger, _ := molog.New(
    molog.WithModuleAsync(asyncmod.Config{
        Routines:    2,
        Buffer:      1024,
        LogOverflow: true,
    }),
)
```

See [asyncModule.md](asyncModule.md) for details.

#### WithModuleRotate(cfg rotatemod.Config) Option

Enables automatic log file rotation with specified configuration.

```go
logger, _ := molog.New(
    molog.WithWriter("file", logFile),
    molog.WithModuleRotate(rotatemod.Config{
        SizeMB:   100,
        AgeHours: 24,
    }),
)
```

See [rotateModule.md](rotateModule.md) for details.

## Runtime Control Methods

### ToggleAsync(enabled bool)

Toggle async mode at runtime.

```go
logger.(*loggerBase).ToggleAsync(true)  // Enable async
logger.(*loggerBase).ToggleAsync(false) // Disable async
```

**Note**: Requires type assertion to access the method.

### AsyncReconfigure(cfg asyncmod.Config) error

Reconfigure async module at runtime.

```go
logger.(*loggerBase).AsyncReconfigure(asyncmod.Config{
    Routines: 4,
    Buffer:   2048,
})
```

### TryRotate()

Forces an immediate rotation check (rotate module). Worker will wake up and perform full sweep.

```go
logger.(*loggerBase).TryRotate()
```

**Note**: Requires type assertion to access the method.

## Internal Architecture

### modulesConfiguration Structure

```go
type modulesConfiguration struct {
    base   base
    async  *asyncmod.Module
    rotate *rotatemod.Module
}
```

This structure contains the nested `base` structure and module pointers. The modular design allows easy addition of new modules.

### base Structure

```go
type base struct {
    writers          map[string]io.Writer
    designatedWriter io.Writer
    handler          handlerType
    addSource        bool
    level            Level
    replaceAttrFuncs map[string]func([]string, slog.Attr) slog.Attr
    customHandler    slog.Handler
    logFiles         []*os.File
}
```

### Internal Methods log() and logAttrs()

All exported logging methods (Debug, Info, Warn, Error, etc.) use internal methods `log()` and `logAttrs()`, which:

1. Check if async module should handle the log (if enabled)
2. Validate context
3. Check if the log level is enabled via `Enabled()`
4. Obtain the program counter (PC) for determining the call source
5. Create a slog.Record and pass it to the handler
6. Check if rotation is needed (if enabled)

This ensures correct file and line number determination for the call.

### IgnorePC Variable

```go
var IgnorePC = false
```

When set to `true`, disables the `runtime.Callers()` call to obtain the program counter. Used for benchmarking overhead.

## Usage Examples

### Basic Usage

```go
package main

import (
    "github.com/Galdoba/molog"
)

func main() {
    logger, _ := molog.New()
    logger.Info("application started")
    logger.Error("connection failed", "error", "timeout")
}
```

### With Options

```go
package main

import (
    "os"
    "github.com/Galdoba/molog"
    "github.com/Galdoba/molog/modules/asyncmod"
)

func main() {
    logger, err := molog.New(
        molog.WithHandlerType(molog.HANDLER_TEXT),
        molog.WithLevel(molog.LevelInfo),
        molog.WithAddSource(true),
        molog.WithWriter("file", os.Stdout),
        molog.WithModuleAsync(asyncmod.Config{
            Routines: 2,
            Buffer:   1024,
        }),
    )
    if err != nil {
        panic(err)
    }

    logger.Info("server listening", "port", 8080)
    logger.Debug("debug info", "data", "test")
}
```

### With Static Attributes

```go
logger, _ := molog.New()
appLogger := logger.With("app", "myapp", "version", "1.0.0")
appLogger.Info("service started")
```

### With Groups

```go
logger, _ := molog.New(molog.WithHandlerType(molog.HANDLER_TEXT))
requestLogger := logger.WithGroup("request")
requestLogger.Info("incoming request", "method", "GET", "path", "/api")
```

## Tests

The project contains a test suite covering the core functionality:

- Creating loggers with various options
- All logging methods (Debug, Info, Warn, Error, Log, LogAttrs)
- Enabled() checking
- With() and WithGroup() methods
- Level and handler type constants
- Working with io.Discard
- Async module functionality

Running tests:

```bash
go test -v ./...
```

## Future Plans

1. **Custom Handlers** - Implementing HANDLER_CUSTOM support
2. **Additional Modules** - See doc/TODO.md for proposed new modules
3. **Extended Features** - Filtering, formatting, additional attributes
