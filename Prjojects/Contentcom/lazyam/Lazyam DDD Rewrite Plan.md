# Lazyam DDD Rewrite Plan

## Domain Model

### 1. Domains (Bounded Contexts)

| Domain | Description |
|--------|-------------|
| **MediaProcessing** | Core domain - transcoding pipeline, script generation |
| **ProjectCatalog** | Amedia metadata management, project/season/episode entities |
| **TaskOrchestration** | Task lifecycle, processing phases, queue management |
| **MediaAnalysis** | File parsing, stream extraction, interlace detection |

---

### 2. Entities (Aggregate Roots)

| Entity | Aggregate | Responsibilities |
|--------|-----------|------------------|
| `Task` | MediaProcessing | Represents a media job, holds processing state, source files |
| `TaskList` | MediaProcessing | Collection of all active tasks |
| `Projects` | ProjectCatalog | Aggregate root for all Amedia projects |
| `AmediaProject` | ProjectCatalog | Movie or series metadata |
| `Season` | ProjectCatalog | Series season |
| `Episode` | ProjectCatalog | Individual episode |
| `MediaProfile` | MediaAnalysis | Technical metadata from ffprobe |
| `Stream` | MediaAnalysis | Video/audio/subtitle stream |
| `Script` | MediaProcessing | Generated transcoding script |

---

### 3. Value Objects

| Value Object | Attributes |
|--------------|------------|
| `FileKey` | Serid, Duration |
| `taskMeta` | GUID, File, TitleOri, TitleRus |
| `Format` | Container metadata (filename, duration, bitrate) |
| `SilenceSegment` | Start, End, Duration, LoudnessBorder |
| `ScriptArgument` | Key, Value |
| `ProcessingPhase` | Phase constant (immutable) |
| `AudioLanguage` | Language code, layout |
| `InterlaceResult` | Frame counts, progressive ratio |
| `TranscodeProfile` | Template type, parameters |

---

### 4. Domain Services

| Service | Domain | Responsibilities |
|---------|--------|------------------|
| `ProcessingService` | MediaProcessing | Orchestrates task processing stages |
| `ScriptGenerationService` | MediaProcessing | Generates FFmpeg scripts from templates |
| `MetadataSyncService` | ProjectCatalog | Synchronizes with external metadata |
| `TaskDiscoveryService` | TaskOrchestration | Scans directories for new tasks |
| `CacheService` | TaskOrchestration | Persists/loads task and project caches |
| `LockService` | TaskOrchestration | Manages processing locks |
| `MediaAnalysisService` | MediaAnalysis | Analyzes files with ffprobe |
| `InterlaceDetectionService` | MediaAnalysis | Runs interlace detection |

---

## Implementation Steps

### Phase 1: Project Structure & Core Domain Setup

1. **Create new module structure**
   ```bash
   mkdir -p cmd/lazyam
   mkdir -p internal/{media,project,orchestration,analysis}/domain
   mkdir -p internal/{media,project,orchestration,analysis}/application
   mkdir -p internal/{media,project,orchestration,analysis}/infrastructure
   mkdir -p pkg/{translit,error}
   ```

2. **Define core domain types**
   - Create `internal/media/domain/entities.go` - Task, SourceFile
   - Create `internal/media/domain/valueobjects.go` - ProcessingPhase, AudioLanguage
   - Create `internal/project/domain/entities.go` - AmediaProject, Season, Episode
   - Create `internal/project/domain/valueobjects.go` - FileKey, ProjectMetadata

3. **Implement Value Objects (Priority: High)**
   - ProcessingPhase enum
   - AudioLanguage struct
   - FileKey value object
   - ScriptArgument struct

---

### Phase 2: Project Catalog Domain

4. **Implement Project domain**
   - `AmediaProject` entity with all metadata fields
   - `Season` and `Episode` entities
   - `Projects` aggregate root
   - Domain methods: SearchByGUID, SearchByFileKey, SeasonEpisode

5. **Implement MetadataSyncService**
   - Read metadata.json from network paths
   - Inject/update projects in aggregate
   - Cache persistence (JSON)

---

### Phase 3: Media Analysis Domain

6. **Implement MediaProfile & Stream**
   - Parse ffprobe JSON output
   - Stream types: video, audio, subtitle
   - Format value object

7. **Implement MediaAnalysisService**
   - ConsumeFile() using ffprobe
   - Validation logic
   - Extract audio languages

8. **Implement InterlaceDetectionService**
   - Run ffprobe idet filter
   - Parse frame counts
   - Calculate progressive ratio

---

### Phase 4: Media Processing Domain

9. **Implement Script Generation**
   - Script entity with template substitution
   - Template registry (Amedia1, Amedia2, Amedia2S, Trailer, ScanInterlace)
   - Argument injection via `|=key=|` syntax

10. **Implement Task entity**
    - Processing state machine
    - Source file collection
    - Methods: AssertReady, ScanSources, AssessInterlace

11. **Implement TaskList aggregate**
    - Task collection management
    - Cache persistence

---

### Phase 5: Orchestration Domain

12. **Implement ProcessingService (Application Layer)**
    - Stage-based processing pipeline
    - Phase transitions
    - Error handling per phase

13. **Implement TaskDiscoveryService**
    - Scan //192.168.31.4/buffer/IN/@AMEDIA_IN/
    - Discover new task directories
    - Create Task entities

14. **Implement LockService**
    - Create/check/remove lock files
    - Auto-cleanup stale locks

15. **Implement CacheService**
    - JSON serialization for Tasks and Projects
    - Load/Save operations

---

### Phase 6: Infrastructure & Integration

16. **Migrate configuration**
   - Move from internal/appmodule/config
   - Create config.yaml with defaults

17. **Migrate translit package**
   - Preserve Russian transliteration logic

18. **Migrate error package**
   - LazyError with Expected/Unexpected classification

19. **Implement CLI (cmd/lazyam)**
   - Use urfave/cli/v3
   - Wire all services together
   - Dependency injection via constructor

---

### Phase 7: Testing & Refinement

20. **Unit tests for domain entities**
   - Value object immutability
   - Entity business rules

21. **Integration tests for services**
   - Metadata sync
   - Script generation
   - Processing pipeline

22. **End-to-end tests**
   - Full processing cycle
   - Cache round-trip

---

## Directory Structure (Target)

```
lazyam/
├── cmd/
│   └── lazyam/
│       └── main.go
├── internal/
│   ├── media/
│   │   ├── domain/
│   │   │   ├── entities.go        # Task, SourceFile
│   │   │   ├── valueobjects.go    # ProcessingPhase, AudioLanguage
│   │   │   └── repository.go      # TaskRepository interface
│   │   ├── application/
│   │   │   └── processing_service.go
│   │   └── infrastructure/
│   │       ├── cache.go
│   │       └── script_generation.go
│   ├── project/
│   │   ├── domain/
│   │   │   ├── entities.go        # AmediaProject, Season, Episode
│   │   │   ├── valueobjects.go    # FileKey, ProjectMetadata
│   │   │   └── repository.go      # ProjectRepository interface
│   │   ├── application/
│   │   │   └── metadata_sync_service.go
│   │   └── infrastructure/
│   │       ├── cache.go
│   │       └── metadata_reader.go
│   ├── orchestration/
│   │   ├── domain/
│   │   │   ├── lock.go
│   │   │   └── processing_cycle.go
│   │   ├── application/
│   │   │   ├── task_discovery.go
│   │   │   └── coordination.go
│   │   └── infrastructure/
│   │       └── lock_file.go
│   └── analysis/
│       ├── domain/
│       │   ├── media_profile.go
│       │   ├── stream.go
│       │   └── interlace_result.go
│       ├── application/
│       │   └── media_analysis_service.go
│       └── infrastructure/
│           └── ffprobe_client.go
├── pkg/
│   ├── translit/
│   │   └── (preserve existing)
│   └── error/
│       └── (preserve existing)
├── config/
│   └── config.yaml
└── go.mod
```

---

## Key Principles

1. **Aggregates boundaries** - Task and Projects are separate aggregates; don't cross-reference internally
2. **Domain Services** - Use services for operations that span multiple entities
3. **Value Objects** - All metadata structures should be immutable
4. **Repository pattern** - Abstract persistence behind interfaces
5. **Dependency inversion** - Infrastructure depends on domain interfaces
6. **No business logic in infrastructure** - ffprobe calls, file I/O in infrastructure layer only
