---
updated_at: 2026-04-22T00:39:49.564+10:00
---
# Architectural Complexity Framework for CLI Applications

**Document Version**: 1.0  
**Purpose**: Provide a shared vocabulary and decision framework for scaling CLI architecture appropriately to project needs.  
**Scope**: Go applications using `urfave/cli/v3` or similar frameworks.

---

## 1. Definition of Architectural Complexity

> **Architectural complexity** is defined as the cardinality of component types and their concrete implementations within a project boundary, where a *component* is any logical unit requiring more than one function to implement (e.g., services, repositories, commands, adapters, domain aggregates).

**Key metrics**:
| Metric | Description |
|--------|-------------|
| **Component Types** | Distinct categories of responsibility (e.g., `Logger`, `Repository`, `Service`) |
| **Implementations per Type** | Number of concrete variants for a given type (e.g., `FileLogger`, `StderrLogger`, `CloudWatchLogger`) |
| **Cross-Cutting Concerns** | Components that interact with multiple layers (e.g., `AuthMiddleware`, `MetricsInterceptor`) |
| **Composition Depth** | Maximum nesting of component dependencies (e.g., `Command → Service → Repository → Adapter`) |

**Formula (heuristic)**:
```
Complexity Score ≈ Σ(Component Types × Implementations per Type) × Composition Depth
```

> This framework is not prescriptive. It is a diagnostic tool to answer: *"Does my architecture match my problem space?"*

---

## 2. Complexity Levels (0–9)

| Level | Term | Description | Component Types | Total Implementations | Example Projects |
|-------|------|-------------|-----------------|----------------------|-----------------|
| **0** | `Script` | Single-file executable; no subcommands; trivial flag parsing; all logic in `main()` or 1–2 helper functions. | 0–1 | 1–2 | `hello.go`, `backup-one-liner.go` |
| **1** | `Utility` | Single command with 3–10 flags; 1–2 components (e.g., config loader, file writer); each component has exactly one implementation. | 1–2 | 2–4 | `jsonfmt`, `grep-lite`, `envsubst-go` |
| **2** | `Tool` | 2–5 subcommands; 3–5 component types; each type has one implementation; minimal cross-cutting concerns. | 3–5 | 4–8 | `taskflow status`, `git-credential-helper` |
| **3** | `Modular CLI` | 5–15 subcommands; 5–10 component types; 1–2 implementations for 2–3 key types (e.g., logger: file+stderr); explicit dependency injection. | 5–10 | 8–20 | Full `taskflow` (TUI/serve/status), `gh` (basic usage) |
| **4** | `Application Suite` | Multiple related CLIs sharing core logic; 10–20 component types; 2–3 implementations for 3–5 types; plugin points for extensions. | 10–20 | 20–50 | `kubectl` (core), `terraform` (basic providers), `docker cli` |
| **5** | `Platform CLI` | CLI as primary interface to a platform; 20–40 component types; 2–4 implementations for most types; runtime component selection; feature flags. | 20–40 | 50–120 | `aws cli`, `gcloud`, `kubectl` + operators |
| **6** | `Ecosystem Gateway` | Orchestrates multiple services/repos; 40–80 component types; 3–5 implementations per type; dynamic discovery; multi-tenant support. | 40–80 | 120–300 | Internal dev tools at mid-size tech company; `helm` + plugins |
| **7** | `Distributed Orchestration Layer` | Entry point to distributed system; 80–150 component types across multiple repos; 5–10 implementations per type; runtime protocol/region selection. | 80–150 | 300–800 | `kubectl` at scale + custom CRDs; service mesh CLI |
| **8** | `Meta-Platform` | Framework that generates/manages other CLIs; 150–300 component types; components are composable frameworks; plugin marketplace; versioned APIs. | 150–300 | 800–2000 | Internal developer platform at FAANG; `buf` + plugin ecosystem |
| **9** | `Sociotechnical System` | CLI embedded in organizational workflow engine; 300+ component types across dozens of repos; 10+ implementations per type; governance/compliance baked into component selection. | 300+ | 2000+ | Google internal tooling; Facebook infrastructure CLI; enterprise service mesh control plane |

---

## 3. Component Taxonomy

For counting purposes, components are categorized as follows:

| Category | Examples | Counting Rule |
|----------|----------|---------------|
| **Commands** | `serve`, `status`, `migrate` | Each `cli.Command` definition = 1 type |
| **Services / Use Cases** | `ProcessManager`, `StateMonitor` | Each business logic unit = 1 type |
| **Repositories** | `StateRepo`, `TaskRepo` | Each data access abstraction = 1 type |
| **Adapters** | `GitHubClient`, `SlackNotifier` | Each external API wrapper = 1 type |
| **Infrastructure** | `Logger`, `ConfigLoader`, `MetricsEmitter` | Each cross-cutting concern = 1 type |
| **Domain** | `Task`, `User`, `Status` | Each aggregate root or value object = 1 type |
| **Policies** | `RetryPolicy`, `AuthStrategy`, `CacheStrategy` | Each swappable behavior = 1 type |

**Implementation counting**:
- `FileLogger` and `StderrLogger` = 2 implementations of type `Logger`
- `PostgresRepo` and `MemoryRepo` = 2 implementations of type `StateRepo`
- `DefaultRetry` and `ExponentialBackoff` = 2 implementations of type `RetryPolicy`

---

## 4. Decision Guidance by Level

### Levels 0–2 (`Script` → `Tool`)
**Recommended approach**:
- Place all code in `main.go` or 1–2 files
- Use `urfave/cli` for flag parsing only
- No interfaces; no dependency injection
- Test via `main()` with table-driven tests

**When to level up**:
- You need a third subcommand with different dependencies
- You find yourself copying flag parsing logic
- Unit testing requires mocking a function with >3 parameters

### Levels 3–4 (`Modular CLI` → `Application Suite`)
**Recommended approach**:
- Adopt the architecture in [[Architectural Recommendations for CLI Applications using `urfave.cli.v3` (v 1.2)|this document]] (builder in `ActionFunc`, interfaces defined by consumer)
- Split `cmd/`, `internal/app/`, `internal/infra/`
- Use functional options for builder flexibility
- Write unit tests for `ActionFunc` with mocked dependencies

**When to level up**:
- You have 3+ CLIs that share >30% of logic
- You need runtime plugin loading
- Component selection depends on config/env at runtime

### Levels 5–7 (`Platform CLI` → `Distributed Orchestration`)
**Recommended approach**:
- Introduce a `pkg/` layer for shared, versioned component interfaces
- Use code generation for boilerplate (e.g., `mockgen`, `stringer`)
- Implement runtime component resolution (e.g., factory pattern with config-driven selection)
- Add observability hooks to every component boundary

**When to level up**:
- You manage components across >3 Git repositories
- Component compatibility is governed by semantic versioning contracts
- You need to A/B test component implementations in production

### Levels 8–9 (`Meta-Platform` → `Sociotechnical System`)
**Recommended approach**:
- Treat components as independently versioned modules
- Implement a component registry with discovery, health checks, and fallback logic
- Embed governance policies (audit, compliance, cost) into component selection
- Use feature flags to control component rollout per tenant/region

**Warning**: At this scale, architecture decisions require cross-team RFCs and formal review. Do not attempt to "refactor up" to this level without dedicated platform engineering support.

---

## 5. Anti-Patterns by Level

| Level | Anti-Pattern | Consequence |
|-------|-------------|-------------|
| 0–1 | Over-engineering with interfaces and DI | Slows development; adds cognitive load without benefit |
| 2–3 | Global state / `init()` for dependencies | Makes testing impossible; hides dependencies |
| 4–5 | Hard-coded component selection | Prevents multi-tenant or multi-cloud support |
| 6–7 | No component versioning or compatibility checks | Causes runtime failures when dependencies update |
| 8–9 | Centralized control of all component logic | Creates single point of failure; slows innovation |

---

## 6. Using This Framework

### For new projects:
1. Estimate your expected complexity level in 12 months.
2. Start with the architecture recommended for that level.
3. Re-evaluate quarterly: if you've crossed a threshold, plan a refactoring sprint.

### For existing projects:
1. Count your component types and implementations (use `golangci-lint run --enable=revive` + custom rules if needed).
2. Plot your current level on the table above.
3. If you are **below** your problem space: plan incremental adoption of patterns from the next level.
4. If you are **above** your problem space: schedule simplification (remove unused abstractions, collapse layers).

### For architecture reviews:
- Ask: *"What complexity level does this change assume? Does it match our current level?"*
- Reject changes that introduce Level-5 patterns into a Level-2 codebase without justification.
- Require a "complexity impact statement" for changes that add new component types.

---

## 7. Relationship to "When to Break the Rules"

This framework enables the future note *"When to Break the Rules"* by providing objective criteria:

> **Rule-breaking is justified when**:
> 1. Your project's complexity level is **lower** than the rule assumes (e.g., no need for builder pattern at Level 1).
> 2. Your project's complexity level is **higher** and the rule creates bottlenecks (e.g., strict layering slows plugin development at Level 7).
> 3. You have **measured evidence** that the rule's cost exceeds its benefit for your use case.

Without a shared definition of complexity, "breaking rules" becomes subjective and contentious. This framework makes the conversation data-driven.

---

## Appendix A: Quick Self-Assessment

Answer these questions to estimate your level:

1. How many subcommands does your CLI have?
   - 1 → Level 0–1
   - 2–5 → Level 2
   - 6–15 → Level 3
   - 16+ → Level 4+

2. How many distinct data stores or external APIs do you integrate with?
   - 0–1 → Level 0–2
   - 2–3 → Level 3–4
   - 4–7 → Level 5–6
   - 8+ → Level 7+

3. Do you need to swap implementations at runtime (e.g., DB driver, auth provider)?
   - No → Level 0–3
   - Yes, via config → Level 4–6
   - Yes, via API/registry → Level 7+

4. How many Git repositories contain code for your CLI?
   - 1 → Level 0–5
   - 2–5 → Level 6–7
   - 6+ → Level 8+

**Scoring**: Take the highest level indicated by any answer. That is your baseline complexity.

---

## Appendix B: Component Counting Example (`taskflow` at Level 3)

```text
Component Types (8):
1. Command (root, serve, status)          → 3 implementations
2. Service (ProcessManager, StateMonitor) → 2 implementations
3. Repository (StateRepo)                 → 2 implementations (LocalFS, RemoteAPI)
4. Logger                                 → 2 implementations (FileOnly, StderrAndFile)
5. ConfigLoader                           → 1 implementation
6. Domain Model (Task, Status)            → 2 types (not counted as "implementations")
7. RetryPolicy                            → 1 implementation
8. MetricsEmitter                         → 1 implementation

Total Implementations: 3+2+2+2+1+1+1 = 12
Composition Depth: Command → Service → Repository → Adapter = 4

Complexity Score ≈ 8 types × 1.5 avg impls × depth 4 = 48 → Level 3 range (8–20 impls)
```

> Note: Domain models are counted as *types* but not *implementations* unless they have swappable variants (e.g., `TaskV1`, `TaskV2`).

---

## Change History

| Version | Date | Changes |
|---------|------|---------|
| 1.0 | 2026 | Initial publication |

---

> **Next Step**: This framework enables the companion note *"When to Break the Rules"* by providing objective criteria for rule applicability. Draft that note next, referencing specific levels from this document.

