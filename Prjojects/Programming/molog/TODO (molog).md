---
updated_at: 2026-03-13T08:13:07.586+10:00
tags:
  - logger
  - modular
---
# New Module Analysis

This document analyzes potential new modules for molog based on the AGENTS.md specifications and user requirements.

## Evaluation Criteria

- **Usefulness**: How much value does it add to users?
- **Feasibility**: How complex is implementation?
- **Dependencies**: What external packages are needed?
- **Priority**: Recommended order of implementation

---

## Candidate Modules

### 1. Filter Module

**Description**: Allow filtering log messages based on level, attributes, or custom predicates.

**Use Cases**:
- Production: Debug/Info messages filtered out
- Development: Only Error messages to file, all to console
- Tenant isolation: Filter by tenant ID in attributes

**Feasibility**: Low - can use existing level filtering + ReplaceAttr

**Proposed API**:
```go
WithFilter(predicate func(Level, string, ...any) bool)
```

**Priority**: Medium

---

### 2. Formatter Module

**Description**: Custom message formatting beyond JSON/TEXT handlers.

**Use Cases**:
- Custom timestamp formats
- Colored output for terminal
- Structured logging with custom keys

**Feasibility**: Medium - requires ReplaceAttr or custom handler

**Proposed API**:
```go
WithFormatter(format FormatType)
FormatType = "json" | "text" | "custom"
WithCustomFormat(fn func(*slog.Record) string)
```

**Priority**: Low (ReplaceAttr already covers most use cases)

---

### 3. Rate Limiter Module

**Description**: Limit log message rate to prevent log flooding.

**Use Cases**:
- Prevent log spam from retry loops
- Protect downstream log services
- Debug bursts with configurable limits

**Feasibility**: Medium - needs token bucket or similar

**Proposed API**:
```go
WithRateLimit(messagesPerSecond int, burst int)
```

**Priority**: High (complements async module)

---

### 4. Hook Module

**Description**: Execute custom callbacks on log events.

**Use Cases**:
- Alert on Error level
- Send to external systems (Sentry, PagerDuty)
- Metrics collection

**Feasibility**: Low - simple callback mechanism

**Proposed API**:
```go
WithHook(level Level, fn func(*slog.Record))
```

**Priority**: High (very useful for production)

---

### 5. Buffer/Flush Module

**Description**: Buffer log messages and flush on condition (size, time, error).

**Use Cases**:
- Batch writes to reduce I/O
- Flush on Error for crash reporting
- Time-based flushing

**Feasibility**: Medium - similar to async but different trigger

**Proposed API**:
```go
WithBuffer(size int, flushOn FlushCondition)
FlushCondition = "full" | "error" | "timeout" | "manual"
Flush()
```

**Priority**: Medium

---

### 6. Context Enrichment Module

**Description**: Automatically enrich logs with context (request ID, user ID, etc.).

**Use Cases**:
- HTTP request tracing
- User session tracking
- Service mesh integration

**Feasibility**: Low - can use With() and WithGroup()

**Proposed API**:
```go
WithContextEnricher(fn func(context.Context) []any)
```

**Priority**: Low (already possible with With())

---

### 7. Multi-Handler Module

**Description**: Route logs to different handlers based on level/attributes.

**Use Cases**:
- Errors to file, info to console
- Different formats for different destinations
- Sensitive data to secure storage

**Feasibility**: Low - can compose handlers

**Proposed API**:
```go
WithRouter(routes []Route)
Route{Level: LevelError, Handler: errorHandler}
```

**Priority**: Medium

---

### 8. Hook into Async Module

**Description**: Add hooks/callbacks within async processing.

**Use Cases**:
- Track async queue depth
- Monitor dropped messages
- Custom retry logic

**Feasibility**: Low - add to existing async module

**Proposed API**:
```go
WithAsyncHook(fn func(event AsyncEvent))
AsyncEvent = "queued" | "processed" | "dropped" | "overflow"
```

**Priority**: Low

---

## Recommended Priority Order

### High Priority (Should Implement Next)

1. **Hook Module** - Most versatile, useful for alerting/metrics
2. **Rate Limiter Module** - Complements async, prevents flooding

### Medium Priority

3. **Buffer/Flush Module** - Good for batch I/O optimization
4. **Multi-Handler Module** - Simple routing functionality
5. **Filter Module** - Basic filtering already possible

### Low Priority

6. **Context Enrichment** - Already achievable with existing API
7. **Formatter Module** - ReplaceAttr covers most cases
8. **Async Hooks** - Nice to have for monitoring

---

## Implementation Notes

### Module Placement Decision

**Option A**: Keep in main package (like async)
- Pros: Simpler API, consistent with current design
- Cons: Package grows large

**Option B**: Subpackage (e.g., `molog/filter`)
- Pros: Better organization, smaller main package
- Cons: More complex API (`molog.New(molog.WithModule(filter.New()))`)

**Recommendation**: Keep in main package for now (as per AGENTS.md guidance), revisit if >5 modules added.

### Common Module Structure

```go
type Module struct {
    // ... module-specific fields
}

func New() *Module { ... }

// Option function pattern
func WithModule(m *Module) Option {
    return func(cfg *modulesConfiguration) error {
        cfg.module = m
        return nil
    }
}
```

---

## Decision Required

Which module should be implemented next?

1. **Hook Module** - Alerts, metrics, external integrations
2. **Rate Limiter** - Prevent flooding, complement async
3. **Buffer/Flush** - Batch I/O optimization
4. **Other** - Suggest a new module
