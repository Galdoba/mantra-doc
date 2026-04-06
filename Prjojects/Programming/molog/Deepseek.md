---
updated_at: 2026-04-03T12:03:47.995+10:00
tags:
  - logger
  - modular
---
Based on a thorough review of your `molog` codebase, I can provide a detailed evaluation and a concrete proposal for a new log rotation module, following the established patterns of your project.

### 1. Codebase Evaluation

Your project is well-structured and follows strong Go conventions. Here are my key observations:

*   **Modular Architecture**: The design is excellent. The `modulesConfiguration` struct in `config.go` and the `Option` pattern in `options.go` provide a clean, extensible foundation for adding new features like a log rotator. The separation of the `asyncm` module into `internal/` is a great precedent.
*   **API Consistency**: The `Logger` interface in `base.go` correctly mirrors `slog.Logger`. The `With*` option functions are consistently implemented, returning an `Option` type.
*   **Runtime Control**: The `control.go` file demonstrates a clear pattern for modules that need runtime management (like `ToggleAsync` and `AsyncReconfigure`). This is a good model for a log rotator, which might need a manual rotation trigger.
*   **Documentation**: The `doc/` directory and `AGENTS.md` are thorough, which is invaluable for maintainability. The `asyncModule.md` provides a perfect blueprint for documenting future modules.
*   **Internal Utility**: The `value.IgnorePC` flag in `internal/value/values.go` shows attention to detail for benchmarking and testing, which is a positive sign.

Overall, the codebase is clean, idiomatic, and ready for the extension you're planning.

### 2. Proposal: The Log Rotation Module (`logrotatem`)

This module would manage log file rotation based on configurable triggers, ensuring log files don't grow unbounded. It would be designed as a non-invasive wrapper around a file writer.

#### Module Parameters (Configuration Options)

Following the `WithAsync` pattern, the main option would be `WithLogRotate`. It would accept functional options to configure its behavior, keeping the API clean and extensible.

**Main Constructor Option:**
```go
// WithLogRotate enables log file rotation for a specified writer.
// It accepts optional configuration functions to fine-tune rotation behavior.
func WithLogRotate(writerKey string, opts ...LogRotateOption) Option
```

**Configuration Options (`LogRotateOption`):**

| Option | Description | Default | Example |
| :--- | :--- | :--- | :--- |
| `WithMaxSizeMB(maxMB int)` | Rotate when file size exceeds this limit (in MB). `0` disables size-based rotation. | `100` MB | `WithMaxSizeMB(50)` |
| `WithMaxAge(d time.Duration)` | Rotate files older than this duration (e.g., `24 * time.Hour`). `0` disables age-based rotation. | `0` (disabled) | `WithMaxAge(7 * 24 * time.Hour)` |
| `WithMaxBackups(count int)` | Number of rotated archives to keep. Older files are deleted. `0` means keep all. | `0` (keep all) | `WithMaxBackups(5)` |
| `WithCompress(compress bool)` | Compress rotated files (e.g., using gzip). | `false` | `WithCompress(true)` |
| `WithFilenamePattern(pattern string)` | Pattern for rotated file names. Supports `{time}` (rotation timestamp) and `{index}`. | `"{original}.{time}.log"` | `WithFilenamePattern("app-{time}.log.gz")` |
| `WithRotateOnStart(rotate bool)` | Rotate immediately when the logger starts (useful for daily logs). | `false` | `WithRotateOnStart(true)` |

#### Decision Point & Triggers

The core decision—*when to rotate*—would be made immediately before a log line is written. This ensures file size is checked against the most recent write.

**Proposed Triggers & Logic:**

1.  **Size Trigger**: The module would track the current file size. Before writing a new log record, it would calculate the estimated new size and compare it against `MaxSizeMB`. If it exceeds the limit, rotation is triggered *before* this record is written, ensuring the new record goes into a fresh file.
2.  **Time Trigger (Age)**: At the start of each write operation, the module would check if the file's creation time (or last rotation time) is older than `MaxAge`. If so, rotation is triggered.
3.  **Manual Trigger**: A runtime method (e.g., `logger.RotateNow(writerKey)`) would allow the application to force a rotation on demand (e.g., in response to a `SIGHUP` signal).

The module would evaluate triggers in a logical order, performing the rotation atomically with respect to the write operation. File locking would be necessary in a concurrent environment to prevent corruption during the rotation process.

### 3. Additional Suggestions & Considerations

*   **Integration with Async Module**: For high-performance applications, consider how the log rotator would interact with the `asyncm` module. A good approach is to have the async worker perform the rotation check and file write, preventing the main application thread from being blocked by file I/O or rotation tasks. The `asyncm.Module` could be extended to accept a "writer middleware" concept.

	*I'm not sure if I want to cross and mix modules beween them selves. rotation should work regardless of async module is on or off. I don't mind Rotate Module working with it's own goroutine/worker in non blocking manner. We must ensure that logger function will not drop logging process if file is busy.*

*   **Atomic File Operations**: Rotation involves multiple steps (close current file, rename, create new). Use filesystem operations like `os.Rename` which are typically atomic within the same mount point. Acquire a lock (e.g., `flock` on Unix) to prevent concurrent processes from rotating simultaneously.
*   **Testing Strategy**: Follow the pattern in `async_test.go`. Use a temporary directory and mock filesystem operations where possible to verify:
    *   Rotation occurs at the correct size.
    *   Time-based rotation works.
    *   Max backup count is respected.
    *   Manual rotation via a `Rotate()` method works.
    *   Concurrency safety.
*   **Documentation**: Create a `doc/logrotate.md` file mirroring `asyncModule.md`. It should detail the parameters, triggers, and usage examples.

This approach builds directly on your existing, well-architected code, adding a powerful feature while maintaining the project's simplicity and modularity.

---

Your concern about keeping modules independent and avoiding cross-dependencies is well-founded. The log rotation functionality should be a self-contained component that works seamlessly whether async mode is on or off, and without risking log loss when the file is busy.

### Proposed Design: Rotation as a Standalone `io.Writer`

The cleanest approach is to implement log rotation as a custom `io.Writer` that can be passed to `WithWriter`. This way, it integrates with the existing writer mechanism and remains completely independent of the async module. The async module (when enabled) will simply pass log records to the handler, which writes to this writer; when async is off, the caller writes directly. In both cases, the rotation logic is handled transparently by the writer.

	Hmmm... making rotation around writer... I see some promlems here, but overall idea seems prommising. We setup writer with rotation ability during setup (compile), than pass it to logger before start to substitute basic io.Writer with our enchanted one. And we never modify it again during runtime... If we want to modify rotation behavour on the fly we shoud reasemble the whole logger again from (for example) config file parameters, on some trigger... time based or by command... Any thoughts on that?
	baseLogger uses io.Writer which is potentialy Multiwriter. Does Go hava a meanings to split Multiwriter - so we do not try to rotate os.Stderr?

#### Internal Architecture of the Rotating Writer

The writer would maintain:

- The current log file handle
- Configuration parameters (max size, max age, max backups, compression flag, filename pattern)
- A mutex to protect the file handle and rotation state
- Optionally, a small buffer for writes that occur during rotation (to avoid blocking)

**Write Operation (`Write(p []byte)`)**  
1. Acquire the mutex.  
2. Check if rotation is needed based on size or age triggers.  
3. If rotation is needed:
   - Close the current file.
   - Rename it to an archive name (using the configured pattern, e.g., `app.log.2026-03-12T15-04-05.gz`).
   - Prune old backups if `maxBackups` is set.
   - Create a new log file.
   - Update internal state (file handle, creation time).
4. Write the data to the new (or existing) file.
5. Release the mutex.

If the file is busy (e.g., during rotation), subsequent `Write` calls will block on the mutex until rotation completes. This guarantees that **no log lines are dropped** – they simply wait a short time. Since rotations are infrequent (e.g., at most once per size threshold or per day), the blocking is negligible.

**Why This Works with Async**  
- **Async enabled**: The write happens in a background worker. The worker’s goroutine will block on the mutex during rotation, but the main application goroutine is unaffected.  
- **Async disabled**: The caller’s goroutine blocks briefly, which is acceptable for most applications given the rarity of rotations.

If you want to eliminate blocking entirely even in synchronous mode, the writer could use a small internal buffer and a dedicated goroutine to perform the actual file I/O (similar to a mini async module). However, that adds complexity and duplicates functionality. I recommend starting with the simpler blocking approach and revisiting only if profiling shows it’s a problem.

#### Configuration and API

Following the pattern of `WithAsync`, you would add:

```go
// WithRotatingWriter enables log rotation for a writer identified by key.
// It accepts functional options to configure rotation behavior.
func WithRotatingWriter(key string, opts ...RotateOption) Option {
    return func(mc *modulesConfiguration) error {
        // Create the rotating writer with given options
        rw := newRotatingWriter(key, opts...)
        // Store it in mc.base.writers under the same key
        mc.base.writers[key] = rw
        // Also keep a reference for runtime control (e.g., manual rotation)
        if mc.rotateWriters == nil {
            mc.rotateWriters = make(map[string]*rotatingWriter)
        }
        mc.rotateWriters[key] = rw
        return nil
    }
}
```

**Rotation Options** (inspired by your `WithAsync` variadic approach but clearer with functional options):

```go
type RotateOption func(*rotatingWriter)

func WithMaxSizeMB(mb int) RotateOption { ... }
func WithMaxAge(d time.Duration) RotateOption { ... }
func WithMaxBackups(n int) RotateOption { ... }
func WithCompress(compress bool) RotateOption { ... }
func WithFilenamePattern(pattern string) RotateOption { ... }
```

**Runtime Control**  
Add a method to `loggerBase` to manually trigger rotation:

```go
func (l *loggerBase) RotateNow(writerKey string) error {
    if l.cfg.rotateWriters == nil {
        return errors.New("no rotating writers configured")
    }
    rw, ok := l.cfg.rotateWriters[writerKey]
    if !ok {
        return fmt.Errorf("writer %q is not a rotating writer", writerKey)
    }
    return rw.rotate() // force rotation now
}
```

	Decidion to pass variatic arguments to optional function was made because Async Module is in "internal" directory. And if I thing User will not be able to call optional functions from internal modeles. I see few decidions:
	- move all optional functions for all modules to dedicated declarative package
	- pass primitives for internal optional functions
	- move modules out of internal directory to give user ability to call modules optional functions directly.
	What do you suggest? Why?


#### Testing Strategy

Follow the pattern in `async_test.go`:
- Use temporary directories and mock files to verify size-based and age-based rotation.
- Test that `maxBackups` correctly removes old files.
- Test compression.
- Test concurrent writes to ensure no data loss or corruption.
- Test manual rotation via `RotateNow`.

#### Documentation

Create `doc/logrotate.md` modeled after `asyncModule.md`, explaining:
- Overview and purpose
- Configuration options with defaults
- How rotation decisions are made (size/time)
- Runtime control (`RotateNow`)
- Interaction with async module (none – they work independently)
- Example usage

### Additional Suggestion: File Locking Across Processes

If multiple processes might write to the same log file (rare, but possible), consider using a system-level lock (e.g., `flock` on Unix) during rotation and writes to prevent corruption. This can be added as an option later if needed.

### Summary

This design keeps the rotation module completely independent, integrates cleanly with your existing writer system, and ensures no log lines are dropped by blocking (or optionally buffering) during rotation. It also provides a clear path for runtime control and testing.


---

Your thoughtful questions touch on important design trade-offs. Let me address each one and then propose a refined, pragmatic design for the log rotation module.

---

### 1. Runtime Modification of Rotation Behavior

**Concern:** If rotation parameters are set only at logger creation, how can we change them later without rebuilding the logger?  

**Observations:**  
- Some applications need dynamic log rotation adjustments (e.g., increasing max size after a config reload).  
- The `async` module already provides a runtime reconfiguration method (`AsyncReconfigure`) that changes worker count and buffer size without rebuilding the logger. This sets a clear precedent.  

**Solution:**  
We can follow the same pattern for rotation. The logger will keep a map of rotating writers keyed by their writer keys. We can then expose methods like:

```go
func (l *loggerBase) SetRotateMaxSize(writerKey string, mb int) error
func (l *loggerBase) SetRotateMaxAge(writerKey string, d time.Duration) error
// etc.
```

These methods would locate the rotating writer and update its internal configuration atomically (protected by a mutex). This allows runtime changes without touching the logger's handler chain.  

**Why this works:**  
- The writer is already stored in `mc.rotateWriters` (as in my earlier snippet).  
- The writer itself holds its configuration and can be modified safely.  
- Subsequent writes will use the new settings.  

**If you prefer not to have runtime modifications**, we could omit these methods and document that rotation settings are immutable after logger creation. But given the precedent of `AsyncReconfigure`, offering similar control seems consistent and valuable.

---

### 2. MultiWriter and Rotating Only Specific Writers

**Concern:** The base logger combines all writers into an `io.MultiWriter`. How do we avoid trying to rotate `os.Stderr`?  

**Answer:** We don't need to split the MultiWriter. The key insight is that each writer in a MultiWriter operates independently.  

- When you call `WithRotatingWriter("file", ...)`, you are adding a rotating writer specifically for the key `"file"`.  
- Other writers (like `os.Stderr` added via `WithWriter("console", os.Stderr)`) remain plain writers.  
- The MultiWriter simply broadcasts each write to **all** writers. The rotating writer rotates its own file; the plain writer writes to stderr unaffected.  

**No special splitting is required.** This is already handled by the existing design: each writer is responsible for its own behavior.  

If you later want to **replace** a writer entirely (e.g., switch from rotating to non-rotating), you would need to rebuild the logger. But runtime modification of rotation parameters on the existing writer is simpler and sufficient for most cases.

---

### 3. API Design: Internal Modules vs. Expressive Options

**Current dilemma:** The `async` module is in `internal/`, so its option `WithAsync` uses variadic `int` parameters. This works but is not self-documenting. For rotation, with multiple parameters, a variadic `int` approach would be confusing and error-prone.

You outlined three possible paths:

- **A.** Move all option functions to a dedicated declarative package (e.g., `molog/options`).  
- **B.** Pass primitives (ints, strings) for internal modules (like `async` does).  
- **C.** Move modules out of `internal/` to expose their types directly.

**My recommendation: Use a functional options pattern that keeps modules internal but provides a clean public API.** This is approach **B with a twist** – not raw primitives, but typed options.

#### How it would work:

1. **Define public option types** in the main `molog` package:

   ```go
   // RotateOption configures a rotating writer.
   type RotateOption func(*rotateConfig)

   type rotateConfig struct {
       maxSizeMB    int
       maxAge       time.Duration
       maxBackups   int
       compress     bool
       filenamePattern string
   }
   ```

2. **Provide public constructor functions** for each option:

   ```go
   func WithMaxSizeMB(mb int) RotateOption {
       return func(c *rotateConfig) { c.maxSizeMB = mb }
   }

   func WithMaxAge(d time.Duration) RotateOption {
       return func(c *rotateConfig) { c.maxAge = d }
   }
   // etc.
   ```

3. **Implement `WithRotatingWriter`** that collects these options and creates the internal writer:

   ```go
   func WithRotatingWriter(key string, opts ...RotateOption) Option {
       return func(mc *modulesConfiguration) error {
           cfg := &rotateConfig{
               maxSizeMB: 100,        // default
               maxAge:    0,           // disabled
               maxBackups: 0,
               compress: false,
               filenamePattern: "{original}.{time}.log",
           }
           for _, opt := range opts {
               opt(cfg)
           }
           // Create the rotating writer (unexported type)
           rw := newRotatingWriter(key, cfg)
           mc.base.writers[key] = rw
           // Store for runtime control
           if mc.rotateWriters == nil {
               mc.rotateWriters = make(map[string]*rotatingWriter)
           }
           mc.rotateWriters[key] = rw
           return nil
       }
   }
   ```

**Advantages of this approach:**

- **Clean, self-documenting API:** Users write `molog.WithMaxSizeMB(50)` instead of remembering argument positions.  
- **Type safety:** Each option is typed; the compiler catches mistakes.  
- **Modules stay internal:** The actual `rotatingWriter` type and its implementation remain in `internal/` (or an unexported subpackage).  
- **Extensible:** New options can be added without breaking existing calls.  
- **Matches `slog` style:** This is exactly how `slog.HandlerOptions` works, so it feels familiar to Go developers.

**Comparison with other options:**

- **Option A (dedicated declarative package):** Would also work, but adds another package for users to import. It might be overkill for a single module. If you later have many modules, a common options package could be justified.  
- **Option C (move out of internal):** Would expose implementation details and increase maintenance burden. It also forces users to import subpackages, making the API more complex.  
- **Option B (primitives only):** Least expressive; fine for simple modules like `async` with 1–2 parameters, but not for rotation.

	Given that we already have one module and planning second. By the end of implementation of Rotate Module we will have **Multiple** modules with Multiple module option types. This means that approaches A and C is clearly justified, in the end there will be 5+ modules. So Option Handling will be clearly it's own thing. I do believe that dedicated public (not internal) options package is a good way to go. Though I might be wrong and you may try to convince me use approach B. Base module may be handled strictly in molog package - so if users not using any modules they will not have to import other packages. But if they using modules - it will have to be imported anyway: one more or less is not really have a meaning when we talk about Multiple. And by design this logger wrapper intended to to be used in small-medium applications with different use cases. Industry level projects should use their own logger fine tuned for particular case.
	Approach C do make design more complex, I but more human-readable. 
	Considering this I incline to have Human-readable Maintainable codebase is very strong point.
	consider these calls:
	```go
	//aproach A
	//is some where in the middle between A and B, but not realy human-readable/maintainable.
	//import 1 package for each module used.
	logger, err := molog.New(
		molog.WithAsync(
			asyncm.WithAsyncParameter1(1),
			asyncm.WithAsyncParameter2("second"),
		),
		molog.WithRotate(
			rotate.WithMaxMB(50),
			rotate.OnStart(true),
		),
	)
	//aproach B
	//very compresed - gives us not bloated codebase, which is faster(?), but require more extensive documentation. Totaly not human-readamle/maintainable
	//no additional import
	logger, err := molog.New(
		molog.WithAsync(4, 1025, 0),
		molog.WithRotate(50, 0, 70),
	)
	//aproach C
	//modules package provide module builders/factories to call individual module in centralized manner and return functions that provide modules.
	//option package provides types and functions for modules builer in modules package
	//individual modules stay internal, 2 packages are imported if we use any modules. 
	logger, err := molog.New(
		modules.Rotate( 
			option.MaxMB(50),
			option.OnStart(50),
		),
	)
	```

What considerations do I miss?
What arguments do you have for each approach?




Given that you already have `async` using primitives, you could **keep it as is** (it's simple) and adopt the functional options approach for rotation, which has more parameters. Consistency across modules is nice, but not mandatory. The user experience for rotation will be much better with named options.

---

### Refined Design Summary

- **Writer-based rotation:** Implement `rotatingWriter` as an `io.Writer` that wraps a file. It checks size/age on each write and rotates if needed.  
- **Configuration:** Use functional options in the public API via `WithRotatingWriter(key, opts...)`.  
- **Runtime control:** Provide methods on `loggerBase` to modify rotation parameters and to force a manual rotation.  
- **MultiWriter compatibility:** No changes needed; rotating and non-rotating writers coexist peacefully.  
- **Module placement:** Keep the rotating writer implementation internal (e.g., `internal/rotatingwriter/`).  

This design gives you the best of both worlds: a clean, extensible API and a well-encapsulated, independent module that works seamlessly with or without the async module.


```go

logger, err := molog.New(
    molog.WithLevel(molog.LevelInfo),
    molog.WithHandlerType(molog.HANDLER_TEXT),
    molog.WithModules(
        module.Async(config.AsyncConfig{
	        Routines:    4,
	        BufferSize:  1024,
	        LogOverflow: true,
	    }),
	    module.Rotate(config.RotateConfig{
	        MaxSizeMB:    50,
	        Compress:     true,
	        RotateOnStart: true,
	    }),
)
```


---

Below is the complete implementation of the `molog` and `module` packages with the async module refactored to use the new design. The code is structured to be clear, maintainable, and free of import cycles.

---

## Package `molog` – Core

### `molog/config.go`

```go
package molog

import (
	"io"
	"log/slog"
	"time"
)

// Level and handler type aliases (unchanged)
type Level = slog.Level
type Attr = slog.Attr
type Record = slog.Record

const (
	LevelDebug Level = slog.LevelDebug
	LevelInfo  Level = slog.LevelInfo
	LevelWarn  Level = slog.LevelWarn
	LevelError Level = slog.LevelError
)

type handlerType string

const (
	HANDLER_JSON   handlerType = "json"
	HANDLER_TEXT   handlerType = "text"
	HANDLER_CUSTOM handlerType = "custom"
)

// AsyncConfig holds all configuration parameters for the async logging module.
// Zero values provide sensible defaults.
type AsyncConfig struct {
	// Routines is the number of worker goroutines. Default: 1
	Routines int

	// BufferSize is the channel buffer size. Default: 0 (unbuffered)
	BufferSize int

	// LogOverflow, when true, logs a warning when the channel buffer overflows.
	// Default: false
	LogOverflow bool
}

// RotateConfig (placeholder for future module)
type RotateConfig struct {
	// ... fields as discussed
}

// ModuleConfigurator is the interface that modules use to configure themselves.
// It is implemented by the unexported modulesConfiguration type.
// This interface lives in the molog package and is the only bridge to the core.
type ModuleConfigurator interface {
	SetAsync(cfg AsyncConfig) error
	SetRotate(cfg RotateConfig) error // for future use
}

// modulesConfiguration holds all module states and implements ModuleConfigurator.
// It is unexported and its fields are only accessible within the molog package.
type modulesConfiguration struct {
	base  base
	async *asyncModule // internal async module implementation
	// rotateWriters map[string]*rotatingWriter // for future
}

// base configuration (unchanged from original)
type base struct {
	writers          map[string]io.Writer
	designatedWriter io.Writer
	handler          handlerType
	addSource        bool
	level            Level
	replaceAttrFuncs map[string]func([]string, slog.Attr) slog.Attr
	customHandler    slog.Handler
}

// SetAsync implements ModuleConfigurator.
func (mc *modulesConfiguration) SetAsync(cfg AsyncConfig) error {
	// Apply defaults
	if cfg.Routines == 0 {
		cfg.Routines = 1
	}
	// BufferSize can be 0 (unbuffered) – that's fine.
	// Create the internal async module.
	mc.async = newAsyncModule(cfg.Routines, cfg.BufferSize, cfg.LogOverflow)
	return nil
}

// SetRotate (placeholder) would create a rotating writer.
func (mc *modulesConfiguration) SetRotate(cfg RotateConfig) error {
	// not implemented yet
	return nil
}
```

### `molog/options.go`

```go
package molog

import (
	"io"
	"log/slog"
)

// Option configures a modulesConfiguration.
type Option func(*modulesConfiguration) error

// WithModules accepts one or more module configuration functions.
// Each function receives a ModuleConfigurator and can call its SetXxx methods.
func WithModules(mods ...func(ModuleConfigurator) error) Option {
	return func(mc *modulesConfiguration) error {
		for _, mod := range mods {
			if err := mod(mc); err != nil {
				return err
			}
		}
		return nil
	}
}

// --- Existing core options (unchanged, just shown for context) ---

func WithWriter(writerKey string, writer io.Writer) Option {
	return func(mc *modulesConfiguration) error {
		if mc.base.writers == nil {
			mc.base.writers = make(map[string]io.Writer)
		}
		if writerKey == "" {
			return errors.New("options: empty key is not allowed")
		}
		delete(mc.base.writers, "")
		mc.base.writers[writerKey] = writer
		return nil
	}
}

func WithHandlerType(ht handlerType) Option {
	return func(mc *modulesConfiguration) error {
		mc.base.handler = ht
		return nil
	}
}

func WithAddSource(addSource bool) Option {
	return func(mc *modulesConfiguration) error {
		mc.base.addSource = addSource
		return nil
	}
}

func WithLevel(level Level) Option {
	return func(mc *modulesConfiguration) error {
		mc.base.level = level
		return nil
	}
}

func WithReplaceAttr(key string, fn func([]string, slog.Attr) slog.Attr) Option {
	return func(mc *modulesConfiguration) error {
		if mc.base.replaceAttrFuncs == nil {
			mc.base.replaceAttrFuncs = make(map[string]func([]string, slog.Attr) slog.Attr)
		}
		mc.base.replaceAttrFuncs[key] = fn
		return nil
	}
}

// WithCustomHandler is kept for future.
func WithCustomHandler(h slog.Handler) Option {
	return func(mc *modulesConfiguration) error {
		mc.base.customHandler = h
		mc.base.handler = HANDLER_CUSTOM
		return nil
	}
}
```

### `molog/async.go` – Internal Async Module

```go
package molog

import (
	"context"
	"log/slog"
	"runtime"
	"sync"
	"time"

	"github.com/Galdoba/molog/internal/value"
)

const asyncRuntimeCallerDepth = 4 // matches previous asyncm depth

// asyncModule is the internal implementation of non‑blocking logging.
// It is completely unexported.
type asyncModule struct {
	enabled     bool
	routines    int
	buffer      int
	logOverflow bool
	mu          sync.Mutex
	recordCh    chan slog.Record
	stopCh      chan struct{}
	wg          sync.WaitGroup
}

// newAsyncModule creates a stopped async module with the given parameters.
func newAsyncModule(routines, buffer int, logOverflow bool) *asyncModule {
	return &asyncModule{
		enabled:     false,
		routines:    routines,
		buffer:      buffer,
		logOverflow: logOverflow,
	}
}

// startWorkers launches the worker goroutines.
func (m *asyncModule) startWorkers(l *slog.Logger) {
	m.recordCh = make(chan slog.Record, m.buffer)
	m.stopCh = make(chan struct{})

	for i := 0; i < m.routines; i++ {
		m.wg.Add(1)
		go m.worker(l)
	}
}

// worker processes log records until stop signal.
func (m *asyncModule) worker(l *slog.Logger) {
	defer m.wg.Done()
	for {
		select {
		case record := <-m.recordCh:
			_ = l.Handler().Handle(context.Background(), record)
		case <-m.stopCh:
			// drain remaining records
			for {
				select {
				case record := <-m.recordCh:
					_ = l.Handler().Handle(context.Background(), record)
				default:
					return
				}
			}
		}
	}
}

// stopWorkers signals workers to stop and waits for them.
func (m *asyncModule) stopWorkers() {
	if m.stopCh != nil {
		close(m.stopCh)
		m.wg.Wait()
		close(m.recordCh)
	}
}

// ToggleEnabled enables or disables async mode at runtime.
func (m *asyncModule) ToggleEnabled(l *slog.Logger, enabled bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if enabled == m.enabled {
		return
	}
	m.enabled = enabled
	if m.enabled {
		m.startWorkers(l)
	} else {
		m.stopWorkers()
	}
}

// Log attempts to send a record asynchronously.
// Returns true if the record was handled (either sent or filtered),
// false if synchronous fallback is needed (e.g., overflow without logging).
func (m *asyncModule) Log(l *slog.Logger, ctx context.Context, level slog.Level, msg string, args ...any) bool {
	if !m.enabled {
		return false
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if !l.Enabled(ctx, level) {
		return true // filtered, no record to send
	}
	var pc uintptr
	if !value.IgnorePC {
		var pcs [1]uintptr
		runtime.Callers(asyncRuntimeCallerDepth, pcs[:])
		pc = pcs[0]
	}
	r := slog.NewRecord(time.Now(), level, msg, pc)
	r.Add(args...)

	select {
	case m.recordCh <- r:
		return true
	default:
		if m.logOverflow {
			l.Warn("log overflow")
			// Returning false will cause synchronous fallback (original logger will handle it).
		}
		return false
	}
}

// LogAttr is the attribute‑based version.
func (m *asyncModule) LogAttr(l *slog.Logger, ctx context.Context, level slog.Level, msg string, attrs ...slog.Attr) bool {
	if !m.enabled {
		return false
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if !l.Enabled(ctx, level) {
		return true
	}
	var pc uintptr
	if !value.IgnorePC {
		var pcs [1]uintptr
		runtime.Callers(asyncRuntimeCallerDepth, pcs[:])
		pc = pcs[0]
	}
	r := slog.NewRecord(time.Now(), level, msg, pc)
	r.AddAttrs(attrs...)

	select {
	case m.recordCh <- r:
		return true
	default:
		if m.logOverflow {
			l.Warn("log overflow")
		}
		return false
	}
}

// Reconfigure changes the module parameters at runtime.
// It stops workers, updates settings, and restarts if enabled.
func (m *asyncModule) Reconfigure(l *slog.Logger, routines, buffer int, logOverflow bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	wasEnabled := m.enabled
	if wasEnabled {
		m.stopWorkers()
	}
	if routines > 0 {
		m.routines = routines
	}
	if buffer >= 0 {
		m.buffer = buffer
	}
	m.logOverflow = logOverflow
	if wasEnabled {
		m.startWorkers(l)
	}
}
```

### `molog/control.go`

```go
package molog

// ToggleAsync enables or disables async mode at runtime.
func (l *loggerBase) ToggleAsync(enabled bool) {
	if l.cfg.async == nil {
		return
	}
	l.cfg.async.ToggleEnabled(l.slog, enabled)
}

// AsyncReconfigure changes the async module parameters at runtime.
// It can modify routines count, buffer size, and overflow logging.
// If a parameter is zero, it keeps the current value.
func (l *loggerBase) AsyncReconfigure(routines, buffer int, logOverflow bool) {
	if l.cfg.async == nil {
		return
	}
	l.cfg.async.Reconfigure(l.slog, routines, buffer, logOverflow)
}
```

### `molog/base.go` (excerpt showing integration)

```go
package molog

// ... (imports, Logger interface, loggerBase struct as before)

func New(opts ...Option) (Logger, error) {
	cfg := modulesConfiguration{
		base: base{
			handler:          HANDLER_JSON,
			addSource:        false,
			level:            LevelDebug,
			replaceAttrFuncs: make(map[string]func([]string, slog.Attr) slog.Attr),
			writers: map[string]io.Writer{
				"": os.Stderr,
			},
		},
	}
	for _, opt := range opts {
		if err := opt(&cfg); err != nil {
			return nil, err
		}
	}

	// Build designated writer from configured writers (unchanged)
	// Build handler (unchanged)
	// ...

	logBase := &loggerBase{slog: slog.New(handler), cfg: cfg}

	// If async module is present and enabled? Actually it starts disabled by default.
	// The user must toggle it on, or we could auto‑start if WithAsync was used?
	// Decision: start disabled; user must call ToggleAsync(true) if they want async.
	// Alternatively, we could start enabled if the module was added. Let's keep consistent with old behavior: added but off.
	// In the old WithAsync, it was started automatically. We'll preserve that: if async module exists, start it.
	if cfg.async != nil {
		cfg.async.ToggleEnabled(logBase.slog, true)
	}
	return logBase, nil
}

// log method now uses async module if present and enabled.
func (l *loggerBase) log(ctx context.Context, level Level, msg string, args ...any) {
	if l.cfg.async != nil && l.cfg.async.Log(l.slog, ctx, level, msg, args...) {
		return // handled asynchronously
	}
	// else synchronous path (unchanged)
	// ...
}

// logAttrs similar.
```

---

## Package `module` – Public Constructors

### `module/async.go`

```go
// Package module provides convenient constructors for molog modules.
package module

import "github.com/Galdoba/molog"

// Async returns a module configuration function that configures the async module.
// It should be passed to molog.WithModules.
func Async(cfg molog.AsyncConfig) func(molog.ModuleConfigurator) error {
	return func(mc molog.ModuleConfigurator) error {
		return mc.SetAsync(cfg)
	}
}

// Rotate (placeholder) would be similar.
func Rotate(cfg molog.RotateConfig) func(molog.ModuleConfigurator) error {
	return func(mc molog.ModuleConfigurator) error {
		return mc.SetRotate(cfg)
	}
}
```

---

## Example Usage

```go
package main

import (
	"github.com/Galdoba/molog"
	"github.com/Galdoba/molog/module"
)

func main() {
	logger, err := molog.New(
		molog.WithLevel(molog.LevelInfo),
		molog.WithHandlerType(molog.HANDLER_TEXT),
		molog.WithModules(
			module.Async(molog.AsyncConfig{
				Routines:    4,
				BufferSize:  1024,
				LogOverflow: true,
			}),
			// module.Rotate(...) would go here when implemented
		),
	)
	if err != nil {
		panic(err)
	}
	logger.Info("hello with async")
	// later, if needed:
	// logger.ToggleAsync(false)
}
```

---

## Why This Works

- **No import cycles**: `molog` does not import `module`. The `module` package imports `molog` only to refer to `AsyncConfig` and `ModuleConfigurator`. The user’s code imports both and ties them together.
- **Encapsulation**: The async implementation (`asyncModule`) is completely unexported and lives inside `molog`. The only public parts are `AsyncConfig` and the `ModuleConfigurator` interface, which is only used by the `module` package.
- **Extensibility**: Adding a new module requires:
  1. Define its config struct in `molog`.
  2. Add a `SetXxx` method to `ModuleConfigurator` and implement it on `modulesConfiguration`.
  3. Add a constructor in the `module` package.
  4. Implement the internal module.
- **Backward compatibility**: Existing code that used `WithAsync` can be migrated gradually; the old `WithAsync` can be deprecated and removed later.

This design achieves the readability of a separate `module` package while keeping the core simple and free of external dependencies.