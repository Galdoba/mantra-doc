---
updated_at: 2026-04-15T19:40:52.520+10:00
tags:
  - package
  - engine
  - domain
---
**Package: `ehex`**  
Path: `ehex`

Provides encoding for extended hexadecimal characters (`0-9`, `A-Z` excluding `I` and `O`), mapping integer values `0–33` to single‑character codes and back. Designed for compact string representation of structured data and use in DTOs.

**Core Types**  
- `Ehex` – Immutable value object combining a code (string), numeric value (`0–33` or negative specials), and optional description.

**Predefined Constants**  
- Standard values obtained via `FromValue` / `FromCode`.  
- Special meta‑values: `Unknown`, `Any`, `Invalid`, `Default`, `Ignore`, `Reserved`, `Masked`, `Extension`, `Placeholder`.

**Key Features**  
- Strict one‑to‑one mapping for standard values `0–33`.  
- Special symbols (`?`, `*`, `!`, `#`, `-`, `&`, `~`, `>`, `.`) represent negative sentinel values.  
- Support for application‑defined extended codes via `New`.  
- `WithDescription` returns a copy, preserving immutability.  
- Implements `fmt.Stringer` (returns code).

**Dependencies**  
- None outside the standard library.

**Test Coverage**  
~100% statement coverage; includes round‑trip tests, special value verification, and zero‑value behaviour.
