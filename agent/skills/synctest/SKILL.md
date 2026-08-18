---
name: synctest
description: testing/synctest golang package usage for deterministic concurrent tests
---

<synctest name="testing/synctest usage">
	<description>
		Package `testing/synctest` helps write deterministic tests for concurrent code: it runs your test in an isolated “bubble” where `time` is faked, and goroutine progress can be driven by waiting for the bubble to become idle.
	</description>
	<basic_pattern>
		```go
		func TestX(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				go func() {
					// your concurrent code
				}()
				// Wait until all other goroutines in the bubble are durably blocked.
				synctest.Wait()
			})
		}
		```
	</basic_pattern>
	<key_behavior>
		1. `synctest.Test(t, f)` runs `f` inside a bubble. Any goroutine started inside is part of the bubble, and `Test` waits for them to exit.
		2. Inside a bubble, `time` uses a fake clock; the initial time is `2000-01-01 00:00:00 UTC`.
		3. Time advances **only when every goroutine in the bubble is durably blocked**.
		4. `synctest.Wait()` blocks until all other goroutines in the bubble are durably blocked (useful to drive goroutines without sleeps/polling).
	</key_behavior>
	<durably_blocked>
		1. What counts as *durably blocked*? It means “blocked and can only be unblocked by another goroutine in the same bubble”.
		2. **Durably blocking operations:**
			1) a blocking send or receive on a channel created within the bubble
			2) a blocking `select` statement where every case is a channel created within the bubble
			3) `sync.Cond.Wait`
			4) `sync.WaitGroup.Wait`, when `WaitGroup.Add` was called within the bubble
			5) `time.Sleep`
		3. **Not durably blocking (common pitfalls):**
			1) locking a `sync.Mutex` or `sync.RWMutex`
			2) I/O (for example, reading from a network socket)
			3) system calls
		</durably_blocked>
		<deadlocks_and_isolation>
			1. When every goroutine in a bubble is durably blocked: `synctest.Wait()` (if called) returns; otherwise time advances to the next event (timer) that will unblock at least one goroutine; if time cannot advance, `synctest.Test` deadlocks and panics.
			2. Channels/timers/tickers created inside a bubble MUST NOT be used from outside the bubble (it panics).
			3. A `sync.WaitGroup` becomes associated with a bubble on the first call to `Add` or `Go`; after that, calling `Add`/`Go` from outside that bubble is a fatal error. A package-level `var wg sync.WaitGroup` cannot be associated with a bubble; use `var wg = new(sync.WaitGroup)` if you need a global.
	</deadlocks_and_isolation>
	<api_restrictions>
		1. `synctest.Test` MUST NOT be called from within a bubble.
		2. `synctest.Wait` MUST NOT be called from outside a bubble, and MUST NOT be called concurrently by multiple goroutines in the same bubble.
		3. Inside `synctest.Test`, you MUST NOT call `t.Run`, `t.Parallel`, or `t.Deadline`.
		4. `t.Cleanup` runs inside the bubble (immediately before `synctest.Test` returns), and `t.Context()` returns a context whose `Done` channel is associated with the bubble.
		5. Cleanup functions and finalizers registered with `runtime.AddCleanup` and `runtime.SetFinalizer` run outside any bubble.
	</api_restrictions>
</synctest>