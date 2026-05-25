# Issues

## 2026-05-25 Using global variables in Goroutines

**What we were trying to do:** For debugging purposes, a global `requestCounter
int32` was added to every request handler, which logged it to `stdout`. Each
additional request incremented the counter by 1.

**What went wrong:** When a load of multiple concurrent users was simulated by
`hey`, the logged numbers were **wayy** out-of-sync & duplicated. It's was
clear it wasn't a reliable counter at all.

`requestCounter += 1` is actually a series of steps:

- reading the variable
- incrementing it
- writing it to memory If multiple goroutines operate at the same time, it can
lead to undefined and corrupt data. This is called a **data race**. **How we
resolved it:** Made incrementing the counter an atomic operation. By using
`sync/atomic` package's `atomic.AddInt32(...)`, the above steps are now one
single atomic operation, and will lead to the correct, desired behaviour.

**What we learned:** When dealing with multi-threading/concurrency, handling
data across threads (goroutines here) is fundamental. Even a operation as
simple as incrementing a integer by one can cause problem if not handled
properly.
