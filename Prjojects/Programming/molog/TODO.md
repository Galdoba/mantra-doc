---
updated_at: 2026-03-15T10:00:00.992+10:00
---
# New Module Analysis

This document analyzes potential new modules for molog based on the AGENTS.md specifications and user requirements.

## Evaluation Criteria

- **Usefulness**: How much value does it add to users?
- **Feasibility**: How complex is implementation?
- **Dependencies**: What external packages are needed?
- **Priority**: Recommended order of implementation

---

## Current Modules

| Module | Status | Setup |
|--------|--------|-------|
| Async | ✅ Stable | `WithModuleAsync(asyncmod.Config{...})` |
| Rotate | 🔧 In Development | `WithModuleRotate(rotatemod.Config{...})` |

---

## Candidate Modules

### 1. Hook Module

**Description**: Execute custom callbacks on log events.

**Use Cases**:
- Alert on Error level
- Send to external systems (Sentry, PagerDuty)
- Metrics collection

**Feasibility**: Low - simple callback mechanism

**Proposed API**:
```go
WithModuleHook(cfg hookmod.Config)
```

**Priority**: High (very useful for production)

---

### 2. Rate Limiter Module

**Description**: Limit log message rate to prevent log flooding.

**Use Cases**:
- Prevent log spam from retry loops
- Protect downstream log services
- Debug bursts with configurable limits

**Feasibility**: Medium - needs token bucket or similar

**Proposed API**:
```go
WithModuleRateLimit(ratelimitmod.Config{MessagesPerSecond: 100, Burst: 200})
```

**Priority**: High (complements async module)

---

### 3. Filter Module

**Description**: Allow filtering log messages based on level, attributes, or custom predicates.

**Use Cases**:
- Production: Debug/Info messages filtered out
- Development: Only Error messages to file, all to console
- Tenant isolation: Filter by tenant ID in attributes

**Feasibility**: Low - can use existing level filtering + ReplaceAttr

**Proposed API**:
```go
WithModuleFilter(filtermod.Config{Predicate: func(...) bool {...}})
```

**Priority**: Medium

---

### 4. Buffer/Flush Module

**Description**: Buffer log messages and flush on condition (size, time, error).

**Use Cases**:
- Batch writes to reduce I/O
- Flush on Error for crash reporting
- Time-based flushing

**Feasibility**: Medium - similar to async but different trigger

**Proposed API**:
```go
WithModuleBuffer(buffermod.Config{Size: 1000, FlushOn: "error"})
```

**Priority**: Medium

---

### 5. Context Enrichment Module

**Description**: Automatically enrich logs with context (request ID, user ID, etc.).

**Use Cases**:
- HTTP request tracing
- User session tracking
- Service mesh integration

**Feasibility**: Low - can use With() and WithGroup()

**Proposed API**:
```go
WithModuleEnrich(enrichmod.Config{Enricher: func(ctx context.Context) []any {...}})
```

**Priority**: Low (already possible with With())

---

## Recommended Priority Order

### High Priority

1. **Hook Module** - Most versatile, useful for alerting/metrics
2. **Rate Limiter Module** - Complements async, prevents flooding

### Medium Priority

3. **Buffer/Flush Module** - Good for batch I/O optimization
4. **Filter Module** - Basic filtering already possible

### Low Priority

5. **Context Enrichment** - Already achievable with existing API
6. **Async Hooks** - Nice to have for monitoring

---

## Module Pattern

All new modules should follow the established pattern:

### Module Constants
Each module must define a `ModuleName` constant:
```go
const ModuleName = "Xxx Module"
```

### Config Structure
```go
type Config struct {
    // Module-specific fields
}

func (cfg Config) Validate() error {
    // Validation logic
    return nil
}
```

### Option Function
```go
func WithModuleXxx(cfg xxxmod.Config) Option {
    return func(mc *modulesConfiguration) error {
        if err := cfg.Validate(); err != nil {
            return fmt.Errorf("invalid %s configuration: %w", xxxmod.ModuleName, err)
        }

        m := xxxmod.New()

        if err := m.Configure(cfg); err != nil {
            return fmt.Errorf("failed to configure %s: %w", xxxmod.ModuleName, err)
        }

        mc.xxx = m
        return nil
    }
}
```

### Module Interface
```go
type Module struct {
    // Module-specific fields
}

func New() *Module { ... }
func (m *Module) Configure(cfg Config) error { ... }
```

---

## Decision Required

Which module should be implemented next?

1. **Hook Module** - Alerts, metrics, external integrations
2. **Rate Limiter** - Prevent flooding, complement async
3. **Buffer/Flush** - Batch I/O optimization
4. **Other** - Suggest a new module
