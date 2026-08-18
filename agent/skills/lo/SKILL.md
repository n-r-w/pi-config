---
name: lo
description: Use this skill when writing, reviewing, refactoring, or explaining Go code that imports github.com/samber/lo or when choosing type-safe functional helpers for slices, maps, strings, tuples, channels, iterators, concurrency utilities, pointer conversion, error-aware transformations, or collection search logic in Go.
---

# samber/lo usage rules for Go

## Supported packages

Use the supported packages as follows:

| Package | Import alias | Use for | Avoid when |
| --- | --- | --- | --- |
| `github.com/samber/lo` | `lo` | Immutable-style helpers for finite slices, maps, strings, tuples, math, channels, search, set operations, type conversion, functions, concurrency utilities, and error handling. | A direct `for` loop or the Go standard library is clearer. |
| `github.com/samber/lo/parallel` | `lop` | Parallel variants of selected slice helpers: `ForEach`, `Map`, `GroupBy`, `PartitionBy`, and `Times`. | The workload needs cancellation, backpressure, bounded workers, or error propagation. |
| `github.com/samber/lo/mutable` | `lom` | Explicit in-place slice operations: `Filter`, `FilterI`, `Map`, `MapI`, `Reverse`, and `Shuffle`. | Callers expect the input slice to remain unchanged. |
| `github.com/samber/lo/it` | `loi` | Lazy iterator helpers for Go iterator workflows. | The collection is already small and eager slice helpers are simpler. |
| `github.com/samber/lo/exp/simd` | `simd` | Experimental SIMD helpers on amd64 with Go 1.26+ and `GOEXPERIMENT=simd`. | Production code needs stable APIs, portable builds, or CPUs without the required SIMD instruction set. |

## Supported helper groups

Use the helper groups as follows:

| Group | Use for | Representative helpers |
| --- | --- | --- |
| Slices | Transforming, filtering, grouping, partitioning, slicing, cloning, chunking, flattening, and ordering finite slices. | `Map`, `Filter`, `FilterMap`, `FlatMap`, `Reduce`, `GroupBy`, `Chunk`, `Window`, `Sliding`, `Take`, `Drop`, `Clone`, `Compact` |
| Maps | Extracting keys or values, filtering, mapping, converting entries, merging, inverting, and turning maps into slices. | `Keys`, `Values`, `HasKey`, `ValueOr`, `PickBy`, `OmitBy`, `Entries`, `FromEntries`, `Assign`, `MapKeys`, `MapValues`, `MapEntries`, `MapToSlice` |
| Math | Numeric ranges and aggregate calculations. | `Range`, `RangeFrom`, `RangeWithSteps`, `Clamp`, `Sum`, `SumBy`, `Product`, `ProductBy`, `Mean`, `MeanBy`, `Mode` |
| Strings | Small string utilities and case conversion. | `RandomString`, `Substring`, `ChunkString`, `RuneLength`, `PascalCase`, `CamelCase`, `KebabCase`, `SnakeCase`, `Words`, `Capitalize`, `Ellipsis` |
| Tuples | Fixed-size tuple construction, unpacking, zipping, unzipping, and cross joins. | `T2` through `T9`, `Zip2` through `Zip9`, `Unzip2` through `Unzip9`, `CrossJoin2` through `CrossJoin9` |
| Time and duration | Duration construction and time extrema. | `Duration`, `Duration0` through `Duration10`, `Earliest`, `Latest`, `EarliestBy`, `LatestBy` |
| Channels | Bridging slices and channels or combining channel flows. | `SliceToChannel`, `ChannelToSlice`, `Buffer`, `BufferWithContext`, `BufferWithTimeout`, `FanIn`, `FanOut` |
| Set and predicate operations | Membership, subset checks, intersections, differences, unions, and removals. | `Contains`, `ContainsBy`, `Every`, `EveryBy`, `Some`, `SomeBy`, `None`, `NoneBy`, `Intersect`, `IntersectBy`, `Difference`, `Union`, `Without`, `WithoutBy`, `WithoutEmpty`, `WithoutNth`, `ElementsMatch`, `ElementsMatchBy` |
| Search | Finding items, indexes, extrema, duplicates, uniques, first, last, nth, and random samples. | `Find`, `FindOrElse`, `FindIndexOf`, `FindLastIndexOf`, `FindKey`, `FindUniques`, `FindDuplicates`, `Min`, `Max`, `First`, `Last`, `Nth`, `Sample`, `Samples` |
| Conditionals | Expression-style condition helpers. | `Ternary`, `TernaryF`, `If`, `ElseIf`, `Else`, `Switch`, `Case`, `Default` |
| Type manipulation | Pointer, nil, empty, any-slice, and coalescing helpers. | `IsNil`, `IsNotNil`, `ToPtr`, `Nil`, `EmptyableToPtr`, `FromPtr`, `FromPtrOr`, `ToSlicePtr`, `FromSlicePtr`, `FromSlicePtrOr`, `ToAnySlice`, `FromAnySlice`, `Empty`, `IsEmpty`, `IsNotEmpty`, `Coalesce`, `CoalesceOrEmpty`, `CoalesceSlice`, `CoalesceSliceOrEmpty`, `CoalesceMap`, `CoalesceMapOrEmpty` |
| Functions | Partial application helpers. | `Partial`, `Partial2`, `Partial3`, `Partial4`, `Partial5` |
| Concurrency | Retry, throttling, debouncing, synchronization, async wrappers, transactions, and wait helpers. | `Attempt`, `AttemptWhile`, `AttemptWithDelay`, `AttemptWhileWithDelay`, `Debounce`, `DebounceBy`, `Throttle`, `ThrottleWithCount`, `ThrottleBy`, `ThrottleByWithCount`, `Synchronize`, `Async`, `Async0` through `Async6`, `Transaction`, `WaitFor`, `WaitForWithContext` |
| Error handling | Panic capture, mandatory values, validation, and assertions. | `Validate`, `Must`, `Try`, `Try1` through `Try6`, `TryOr`, `TryOr1` through `TryOr6`, `TryCatch`, `TryWithErrorValue`, `TryCatchWithErrorValue`, `ErrorsAs`, `Assert`, `Assertf` |
| Constraints | Generic constraints used by helper signatures. | `Clonable` |

## General decision rules

1. Prefer plain Go loops when they are shorter, clearer, need early `break` or `continue`, mutate several variables, or require nuanced control flow.
2. Prefer `lo` helpers when they remove repetitive collection plumbing and keep the transformation obvious.
3. Prefer the Go standard library when `slices`, `maps`, `cmp`, `iter`, `strings`, or `sort` express the same operation clearly.
4. Do not dot-import `github.com/samber/lo`. Always keep the `lo.`, `lop.`, `lom.`, or `loi.` qualifier visible.
5. Keep callbacks pure unless the helper is explicitly used for side effects, such as `ForEach`, `ForEachWhile`, `Attempt`, or `Transaction`.
6. Avoid long chains of nested helpers. Use named intermediate variables when it improves debugging, allocation visibility, or error context.
7. Be explicit about allocation behavior. Most root `lo` collection helpers return new collections; `mutable` helpers modify the input slice.
8. Do not use deprecated root `lo.Reverse` or `lo.Shuffle` in new code. Use `lom.Reverse` or `lom.Shuffle` when in-place mutation is intended.
9. Treat `exp/` packages as unstable. Do not use them in shared libraries or public APIs unless the project explicitly accepts experimental dependencies.

## Slice rules

1. Use `Map` for one-to-one transformations:
```go
names := lo.Map(users, func(u User, _ int) string {
    return u.Name
})
```
2. Use `Filter` or `Reject` for predicate-based selection:
```go
active := lo.Filter(users, func(u User, _ int) bool {
    return u.Active
})
```
3. Use `FilterMap` only when filtering and transforming are naturally one operation:
```go
ids := lo.FilterMap(rows, func(r Row, _ int) (int64, bool) {
    if r.ID == 0 {
        return 0, false
    }
    return r.ID, true
})
```
4. Use `Reduce` when the result is not naturally another slice or map:
```go
total := lo.Reduce(items, func(sum int64, item Item, _ int) int64 {
    return sum + item.Price
}, int64(0))
```
5. Use `GroupBy`, `GroupByMap`, `KeyBy`, `SliceToMap`, or `Associate` for indexing and grouping. Make duplicate-key behavior explicit in review comments or code names when later values overwrite earlier values.
6. Use `Chunk`, `Window`, and `Sliding` only when their edge behavior is acceptable. Prefer explicit code for business-critical batching rules.
7. Use `Clone` before passing a slice to code that may mutate it.
8. Use `Compact` only when the zero value should be removed. Do not use it if zero is a valid business value.

## Error-aware helper rules

1. Use `*Err` variants when the callback can fail and the helper provides a matching error-aware API:
```go
out, err := lo.MapErr(values, func(v string, _ int) (int, error) {
    return strconv.Atoi(v)
})
```
2. Prefer `MapErr`, `FlatMapErr`, `FilterErr`, `GroupByErr`, `GroupByMapErr`, `ReduceErr`, `ReduceRightErr`, `MapEntriesErr`, `MapKeysErr`, `MapValuesErr`, and `MapToSliceErr` over embedding error state in an outer variable.
3. Do not ignore partial results after an error unless the helper documents that partial results are meaningful. Treat returned collections from failed `*Err` calls as invalid.
4. Preserve error context inside callbacks when the item identity matters:
```go
out, err := lo.MapErr(rows, func(row Row, _ int) (User, error) {
    user, err := parseUser(row)
    if err != nil {
        return User{}, fmt.Errorf("parse user %q: %w", row.ID, err)
    }
    return user, nil
})
```
5. Use plain loops when error handling requires rollback, retries per item, metrics per failure, multiple accumulated errors, or cancellation.

## Map rules

1. Use `Keys` and `Values` only when order is irrelevant. Go map iteration order is not stable.
2. Use `PickBy`, `PickByKeys`, `PickByValues`, `OmitBy`, `OmitByKeys`, and `OmitByValues` to express map filtering without manual allocation boilerplate.
3. Use `MapKeys`, `MapValues`, or `MapEntries` for structural map transformations. Verify that transformed keys cannot collide unless overwrite semantics are intended.
4. Use `Entries` and `FromEntries` at API boundaries that naturally represent key-value pairs as slices.
5. Use `Assign` for shallow map merge only. Do not use it for deep merge semantics.
6. Use `ValueOr` for simple fallback lookup; use the native comma-ok form when absence must be distinguished from a stored zero value.

## Search and set rules

1. Use `Contains`, `Every`, `Some`, and `None` for small, readable predicate checks.
2. Use `ContainsBy`, `EveryBy`, `SomeBy`, and `NoneBy` for predicate-based checks over structs or non-comparable conditions.
3. Use `Find`, `FindOrElse`, `First`, `FirstOr`, `FirstOrEmpty`, `Last`, `LastOr`, `LastOrEmpty`, `Nth`, `NthOr`, and `NthOrEmpty` instead of ad hoc boundary checks when the fallback semantics are obvious.
4. Use `Intersect`, `Difference`, `Union`, `Without`, and `WithoutEmpty` for set-like operations. Prefer explicit map-based code for very large collections when allocation and complexity must be tuned.
5. Use `ElementsMatch` when order must not matter. Use `reflect.DeepEqual`, `slices.Equal`, or a direct loop when order must matter.

## Pointer, nil, and empty rules

1. Use `ToPtr` and `FromPtrOr` at boundaries where pointer fields are unavoidable, such as generated API models, optional config, or database adapters.
2. Prefer explicit struct fields over `map[string]any` plus `FromAnySlice` when the schema is known.
3. Use `IsNil` and `IsNotNil` carefully with interfaces. Prefer direct nil checks when the static type makes them reliable.
4. Use `Coalesce`, `CoalesceOrEmpty`, `CoalesceSlice`, `CoalesceSliceOrEmpty`, `CoalesceMap`, and `CoalesceMapOrEmpty` only when the first non-empty value is truly the intended rule.
5. Do not use `Empty[T]()` to hide domain defaults. Use named constants or constructors when the zero value has business meaning.

## Parallel and concurrency rules

1. Use `lop.Map`, `lop.ForEach`, `lop.GroupBy`, `lop.PartitionBy`, and `lop.Times` only for independent CPU-bound or latency-bound work where callback order does not affect correctness.
2. Do not use `parallel` helpers when callbacks mutate shared state unless synchronization is explicit and reviewed.
3. Do not use `parallel` helpers for unbounded collections or expensive external calls when concurrency limits are required. Use a worker pool, `errgroup`, channels, or explicit semaphores.
4. Prefer standard `context.Context` cancellation over `parallel` helpers when work must stop early.
5. Use `Attempt`, `AttemptWithDelay`, `AttemptWhile`, and `AttemptWhileWithDelay` for small retry loops. Use a dedicated retry package or explicit code when jitter, backoff policy, context cancellation, observability, or idempotency is important.
6. Use `Debounce` and `Throttle` only when lifecycle ownership is clear. Avoid package-level debounce or throttle state in tests and request-scoped code.
7. Use `Async` helpers only for simple future-like wrappers. Prefer goroutines, channels, `sync.WaitGroup`, or `errgroup` for structured concurrent workflows.

## Mutable rules

1. Import mutable helpers as `lom`:
```go
import lom "github.com/samber/lo/mutable"
```
2. Use `lom.Filter` and `lom.Map` only when reusing the backing array is intentional and safe:
```go
items = lom.Filter(items, func(item Item) bool {
    return item.Enabled
})
```
3. Never pass a slice to `lom` helpers when other goroutines or callers retain aliases to the same backing array.
4. Prefer root `lo` helpers in public APIs, tests, and review examples unless mutation is the point being demonstrated.
5. Treat `lom.Reverse` and `lom.Shuffle` as in-place operations; clone first if the original order is still needed.

## Iterator rules

1. Import iterator helpers as `loi`:
```go
import loi "github.com/samber/lo/it"
```
2. Use `loi` when the input is already an iterator or when lazy composition avoids materializing intermediate slices.
3. Use root `lo` helpers when the data is already a slice and eager allocation is acceptable.
4. Keep iterator pipelines short. Convert to a named slice or map before complex branching, logging, or error handling.
5. Do not mix `loi` and channel helpers unless the ownership and termination behavior are explicit.

## Review checklist

1. A plain loop would not be clearer than the selected `lo` helper.
2. Standard-library `slices`, `maps`, `strings`, `sort`, `cmp`, or `iter` would not express the same operation more clearly.
3. Callback side effects are absent or intentional.
4. Returned collection allocation is acceptable for the hot path.
5. Map key collision and map iteration ordering are handled explicitly where relevant.
6. `*Err` variants are used instead of external error variables inside callbacks.
7. `parallel` helpers are not used where cancellation, bounded concurrency, ordering side effects, or shared mutation matter.
8. `mutable` helpers are not used on aliased slices unless mutation is intentional.
9. Deprecated root helpers such as `lo.Reverse` and `lo.Shuffle` are not used in new code.
10. Experimental `exp/` packages do not leak into stable APIs.
