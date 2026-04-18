---
updated_at: 2026-04-15T17:24:46.924+10:00
tags:
  - package
  - engine
  - domain
---
**Package: `dice`**  
Path: `cepheus/internal/domain/engine/dice`

Provides a flexible dice rolling system with expression parsing, modifier chaining, and support for common RPG mechanics (D66, Flux, Variance).

**Core Types**  
- `Manager` – Coordinates dice rolling, expression caching, and result interpretation. Safe for concurrent use.  
- `Result` – Contains the rolled dice and raw values; returns copies to prevent mutation.  
- `Roller` – Interface for custom die‑rolling implementations (default uses `math/rand` with optional seed).

**Key Features**  
- Parses standard dice notation (e.g., `3d6+2`, `4d6:dl1:2e`).  
- Modifier system with priority ordering: drop lowest/highest, add to each die, add to individual dice, multiply/divide, sum, and simple additive.  
- Caches parsed expressions for performance.  
- Deterministic seeding via `New(seed string)` for reproducible rolls.  
- Specialized methods: `D66` (two‑digit percentile), `Flux` (d6‑d6), `FluxGood`/`FluxBad`, and `Variance` (0.0–1.0).  
- Package‑level functions (`Roll`, `MustRoll`, `D66`, etc.) use a thread‑safe default manager.

**Dependencies**  
- Standard library only (`fmt`, `math/rand`, `sync`, `slices`, `strconv`, `strings`, `time`).

**Test Coverage**  
~82% statement coverage; includes unit tests for parsing, deterministic rolling, modifier ranges, concurrency safety, and edge cases.