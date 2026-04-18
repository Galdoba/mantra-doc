---
updated_at: 2026-04-15T17:35:54.141+10:00
tags:
  - package
  - engine
  - domain
---
**Package: `tables`**  
Path: `cepheus/internal/domain/engine/tables`

Provides data structures and methods for defining and rolling on system agnostic random generation tables, commonly used in tabletop RPGs and procedural content generation.

**Core Types**  
- `GameTable` – Represents a single random table with a name, dice expression (or D66 flag), and a map of result ranges to outcomes. Includes comprehensive validation (no holes, bounds checks).  
- `Collection` – Groups multiple named tables, supports rolling on a table (`Roll`) and cascading through result chains (`RollCascade`). Records roll history.

**Key Features**  
- Supports standard [[dice]] notation (e.g., `2d6+1`) via the `TableRoller` interface.  
- Handles D66 tables (2d6 percentile‑style, e.g., `"11"`–`"66"`).  
- Flexible index notation: `"2-5"`, `"8+"`, `"0-"`, comma‑separated lists.  
- Strict validation ensures tables are complete and ready for use.  
- Cascading rolls allow results to reference other tables, creating multi‑step generation chains (with depth limit protection).

**Dependencies**  
- Relies on an implementation of `TableRoller` (e.g., the internal `dice` package) for actual dice rolling.

**Test Coverage**  
~95% statement coverage; includes unit tests for all major functionality and edge cases.

**Examples**

```json
{
  "name": "Table Name",
  "expression": "2d10",
  "data": {
	"0-": "None",
	"1": "0 + 1d10-1 degrees",
    "2 - 5": "0 + 1d10-1 degrees",
    "6 - 9": "10 + 1d10-1 degrees",
    "10 - 13": "20 + 1d10-1 degrees",
    "14 - 17": "30 + 1d10-1 degrees",
    "18 - 19": "40 + 1d10-1 degrees",
    "20+": "Extreme"
  },
  "d_66": false // can be omited if false
}
```