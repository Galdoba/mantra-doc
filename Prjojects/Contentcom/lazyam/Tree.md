# Lazyam DDD Repository Tree

```
lazyam/
├── cmd/
│   └── lazyam/
│       └── main.go                  # Application entry point
├── internal/
│   ├── media/
│   │   ├── domain/
│   │   │   ├── entities.go           # Task, SourceFile
│   │   │   ├── valueobjects.go      # ProcessingPhase, AudioLanguage
│   │   │   └── repository.go        # TaskRepository interface
│   │   ├── application/
│   │   │   └── processing_service.go
│   │   └── infrastructure/
│   │       ├── cache.go             # Task cache persistence
│   │       └── script_generation.go # FFmpeg script generation
│   ├── project/
│   │   ├── domain/
│   │   │   ├── entities.go          # AmediaProject, Season, Episode
│   │   │   ├── valueobjects.go      # FileKey, ProjectMetadata
│   │   │   └── repository.go        # ProjectRepository interface
│   │   ├── application/
│   │   │   └── metadata_sync_service.go
│   │   └── infrastructure/
│   │       ├── cache.go             # Project cache persistence
│   │       └── metadata_reader.go   # Read metadata.json
│   ├── orchestration/
│   │   ├── domain/
│   │   │   ├── lock.go              # Lock entity
│   │   │   └── processing_cycle.go  # Processing cycle entity
│   │   ├── application/
│   │   │   ├── task_discovery.go    # Scan directories for tasks
│   │   │   └── coordination.go       # Task coordination
│   │   └── infrastructure/
│   │       └── lock_file.go         # File-based locking
│   └── analysis/
│       ├── domain/
│       │   ├── media_profile.go     # MediaProfile entity
│       │   ├── stream.go            # Stream entity
│       │   └── interlace_result.go  # InterlaceResult value object
│       ├── application/
│       │   └── media_analysis_service.go
│       └── infrastructure/
│           └── ffprobe_client.go    # FFprobe integration
├── pkg/
│   ├── translit/
│   │   └── (preserved from legacy)  # Russian transliteration
│   └── error/
│       └── (preserved from legacy)  # LazyError types
├── config/
│   └── config.yaml                  # Application configuration
├── go.mod
└── go.sum
```

## Domain Layer Summary

| Bounded Context | Entities | Value Objects | Services |
|-----------------|----------|---------------|----------|
| **MediaProcessing** | Task, TaskList, Script | ProcessingPhase, ScriptArgument | ProcessingService, ScriptGenerationService |
| **ProjectCatalog** | Projects, AmediaProject, Season, Episode | FileKey, taskMeta, Format | MetadataSyncService |
| **TaskOrchestration** | ProcessingCycle | ProcessingPhase | TaskDiscoveryService, CacheService, LockService |
| **MediaAnalysis** | MediaProfile, Stream | AudioLanguage, InterlaceResult, TranscodeProfile | MediaAnalysisService, InterlaceDetectionService |

## Bounded Contexts - Design Decisions

### 1. MediaProcessing (Core Domain)

**Why separate?** This is the core business domain - the reason the application exists. It handles the transcoding pipeline end-to-end, from task creation through script generation.

**Design Decisions:**
- **Aggregate Root: Task** - Represents a single media processing job. Contains processing state, source files, and references to generated scripts. Task must be the root because all other entities (SourceFile, Script) only exist in the context of a Task.
- **TaskList** - Collection aggregate for managing multiple tasks. Provides operations like finding tasks by state, filtering by phase.
- **ProcessingPhase as Value Object** - Immutable enum representing phases (Scan, Analyze, Transcode, etc.). Phases define the state machine transitions and can't exist independently.
- **ScriptGenerationService** - Domain service because script generation requires understanding both Task state and TranscodeProfile parameters. It orchestrates multiple entities without being part of any single one.

**Boundaries:** Does NOT own project metadata - tasks reference projects by ID only. Does NOT directly perform media analysis - delegates to MediaAnalysis domain.

---

### 2. ProjectCatalog (Supporting Domain)

**Why separate?** Manages Amedia metadata that exists independently of processing tasks. A project can exist without any active tasks, and the same project may be processed multiple times (releases, remuxes).

**Design Decisions:**
- **Aggregate Root: Projects** - Collection aggregate that owns all AmediaProject entities. Provides global search operations (by GUID, by FileKey).
- **AmediaProject as Entity** - Has identity (SERID) and lifecycle. Can contain Seasons and Episodes as child entities.
- **Season/Episode as Entities** - Part of AmediaProject aggregate - they don't have independent identity outside the project context.
- **FileKey as Value Object** - Immutable composite of Serid + Duration. Used for matching source files to projects. Two FileKeys with same values should be equal.
- **MetadataSyncService** - Reads external metadata.json and updates the aggregate. This is a domain service because it coordinates between external data source and internal entities.

**Boundaries:** Does NOT handle processing logic - that's MediaProcessing. Does NOT know about file analysis details - only metadata.

---

### 3. TaskOrchestration (Supporting Domain)

**Why separate?** Handles cross-cutting concerns around task lifecycle management, discovery, and coordination. These operations don't belong to any single aggregate but affect multiple ones.

**Design Decisions:**
- **ProcessingCycle as Entity** - Represents one full processing iteration. Manages the overall flow from task discovery through completion.
- **Lock as Entity** - Prevents concurrent processing of the same resource. Must be independent because locks can exist across aggregate boundaries.
- **TaskDiscoveryService** - Scans file system for new tasks. Creates Task entities in MediaProcessing domain but doesn't own them.
- **CacheService** - Persistence abstraction. Lives here because it handles both Task and Project caches.
- **LockService** - Manages lock lifecycle. Infrastructure-aware service for file-based locking.

**Boundaries:** Does NOT define what a Task IS (that's MediaProcessing). Does NOT define what a Project IS (that's ProjectCatalog). Orchestration coordinates these domains.

---

### 4. MediaAnalysis (Generic Subdomain)

**Why separate?** This domain deals with technical analysis of media files using ffprobe. It's a "generic" subdomain - the concepts (streams, codecs, interlace detection) could be reused by other applications. Isolating it allows replacing ffprobe with another tool without affecting core business logic.

**Design Decisions:**
- **MediaProfile as Entity** - Represents technical metadata extracted from ffprobe. Has identity based on the analyzed file path + modification time.
- **Stream as Entity** - Part of MediaProfile aggregate. Represents video/audio/subtitle streams.
- **InterlaceResult as Value Object** - Immutable result of interlace detection. Contains frame counts and calculated progressive ratio.
- **AudioLanguage as Value Object** - Immutable representation of audio language with layout info.
- **MediaAnalysisService** - Executes ffprobe and parses JSON output. Infrastructure-dependent - lives in application layer.
- **InterlaceDetectionService** - Specialized analysis using ffprobe idet filter. Returns InterlaceResult value object.

**Boundaries:** Does NOT make business decisions based on analysis - that's MediaProcessing. Does NOT manage project metadata - that's ProjectCatalog.

---

## Cross-Cutting Design Principles

### Aggregate Boundaries
- **Task** and **Projects** are separate aggregates - they have different lifecycles and can be modified independently
- **No cross-aggregate references** - tasks reference projects by ID, not object references
- **Eventual consistency** between aggregates - task may reference a project that was updated

### Value Object Usage
- All metadata structures are immutable after creation
- ProcessingPhase, FileKey, AudioLanguage, InterlaceResult are pure value types
- Value objects can be freely passed between domains without side effects

### Service Layering
- **Domain Services** - Pure business logic, no infrastructure dependencies
- **Application Services** - Orchestrate domain operations, may use infrastructure
- **Infrastructure Services** - External integrations (ffprobe, file system, network)

### Repository Pattern
- Each aggregate has a repository interface in domain layer
- Infrastructure implements repositories (JSON file storage)
- Application layer depends on domain interfaces (dependency inversion)
