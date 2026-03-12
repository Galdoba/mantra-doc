---
updated_at: 2026-03-13T08:13:34.411+10:00
tags:
  - logger
  - modular
---
# molog - MOdular LOGger

Molog is a modular wrapper around the standard library's `log/slog` package for Go. The library provides a convenient and flexible logger constructor with various configuration options.

## Overview

The project provides full coverage of `*slog.Logger` functionality with extended configuration capabilities and an extensible module system.

## Modules

| Module | Status | Description |
|--------|--------|-------------|
| [base](base.md) | ✅ Stable | Core logger with full slog compatibility |
| [async](asyncModule.md) | ✅ Prototype | Non-blocking async logging |

## Quick Links

- [README](../README.md) - Installation and quick start
- [Base Module](base.md) - Core functionality
- [Async Module](asyncModule.md) - Non-blocking logging