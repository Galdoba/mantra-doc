---
updated_at: 2026-03-15T09:58:26.011+10:00
tags:
  - logging
  - logger
  - modular
---
# molog - MOdular LOGger

Molog is a modular wrapper around the standard library's `log/slog` package for Go. The library provides a convenient and flexible logger constructor with various configuration options.

## Overview

The project provides full coverage of `*slog.Logger` functionality with extended configuration capabilities and an extensible module system.

## Configuration Pattern

Molog uses functional options for configuration:

- **Base options** (With*): Configure core logger settings
- **Module options** (WithModule*): Enable and configure additional modules

```go
logger, err := molog.New(
    molog.WithWriter("console", os.Stderr),
    molog.WithHandlerType(molog.HANDLER_TEXT),
    molog.WithLevel(molog.LevelInfo),
    molog.WithModuleAsync(asyncmod.Config{Routines: 2, Buffer: 1024}),
)
```

## Modules

| Module | Status | Description |
|--------|--------|-------------|
| [base](base.md) | ✅ Stable | Core logger with full slog compatibility |
| [async](asyncModule.md) | ✅ Stable | Non-blocking async logging |
| [rotate](rotateModule.md) | 🔧 In Development | Automatic log file rotation |

## Quick Links

- [README](../README.md) - Installation and quick start
- [Base Module](base.md) - Core functionality
- [Async Module](asyncModule.md) - Non-blocking logging
- [Rotate Module](rotateModule.md) - Log file rotation
