---
updated_at: 2026-03-13T09:58:00.868+10:00
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
├── base.go        # Logger interface, implementation, constructors
├── config.go      # Types, constants, configuration structures
├── options.go     # Option functions for configuration
├── control.go     # Runtime control methods (ToggleAsync, AsyncReconfigure)
└── base_test.go   # Tests
```

## Current Modules

| Module | Status | Description |
|--------|--------|-------------|
| base | ✅ Stable | Core logger functionality |
| async | ✅ Prototype | Non-blocking async logging |

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

### Constructors

| Function | Description |
|----------|-------------|
| `New(opts ...Option) (Logger, error)` | Create logger with options |
| `Default() (Logger, error)` | Get wrapper around slog.Default() |
| `NewWithHandler(h slog.Handler) *loggerBase` | Create logger with custom handler |

### Available Types and Aliases

```go
type Level = slog.Level      // Log level
type Attr = slog.Attr        // Record attribute
type Record = slog.Record    // Log record

const (
    LevelDebug Level = slog.LevelDebug  // Debug level
    LevelInfo  Level = slog.LevelInfo    // Info level
    LevelWarn  Level = slog.LevelWarn    // Warning level
    LevelError Level = slog.LevelError    // Error level
)

type handlerType string

const (
    HANDLER_JSON   handlerType = "json"   // JSON output format
    HANDLER_TEXT   handlerType = "text"   // Text output format
    HANDLER_CUSTOM handlerType = "custom" // Custom handler (not implemented)
)
```

## Constructor New() Options

### WithWriter(writerKey string, writer io.Writer) Option

Adds a writer for log output. The key cannot be empty.

```go
logger, _ := molog.New(
    molog.WithWriter("console", os.Stderr),
)
```

When multiple writers are specified, they are all combined via `io.MultiWriter`.

### WithHandlerType(ht handlerType) Option

Sets the handler type (JSON or TEXT). Default is JSON.

```go
logger, _ := molog.New(
    molog.WithHandlerType(molog.HANDLER_TEXT),
)
```

### WithAddSource(addSource bool) Option

Enables/disables adding source information (file:line). Default is false.

```go
logger, _ := molog.New(
    molog.WithAddSource(true),
)
```

### WithLevel(level Level) Option

Sets the minimum log level. Default is LevelDebug.

```go
logger, _ := molog.New(
    molog.WithLevel(molog.LevelInfo),
)
```

### WithReplaceAttr(key string, fn func([]string, slog.Attr) slog.Attr) Option

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

### WithCustomHandler(h slog.Handler) Option

Allows passing a completely custom handler. **Note: Currently not implemented.**

```go
// Currently not available
```

## Internal Architecture

### modulesConfiguration Structure

```go
type modulesConfiguration struct {
    base  base
    async *asyncm.Module
}
```

This structure contains the nested `base` structure and the async module. The modular design allows easy addition of new modules.

### base Structure

```go
type base struct {
    writers              map[string]io.Writer
    designatedWriter    io.Writer
    handler             handlerType
    addSource           bool
    level               Level
    replaceAttrFuncs    map[string]func([]string, slog.Attr) slog.Attr
    customHandler       slog.Handler
}
```

### Internal Methods log() and logAttrs()

All exported logging methods (Debug, Info, Warn, Error, etc.) use internal methods `log()` and `logAttrs()`, which:

1. Validate context
2. Check if the log level is enabled via `Enabled()`
3. Obtain the program counter (PC) for determining the call source
4. Create a slog.Record and pass it to the handler

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
)

func main() {
    logger, err := molog.New(
        molog.WithHandlerType(molog.HANDLER_TEXT),
        molog.WithLevel(molog.LevelInfo),
        molog.WithAddSource(true),
        molog.WithWriter("file", os.Stdout),
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

Running tests:

```bash
go test -v ./...
```

## Future Plans

1. **Custom Handlers** - Implementing HANDLER_CUSTOM support
2. **Additional Modules** - See doc/TODO.md for proposed new modules
3. **Extended Features** - Filtering, formatting, additional attributes
