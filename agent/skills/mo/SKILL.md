---
name: mo
description: Use this skill when writing, reviewing, refactoring, or explaining Go code that imports github.com/samber/mo or when choosing monadic abstractions for nullable values, fallible computations, discriminated unions, lazy effects, async computations, or stateful computations in Go.
---

# samber/mo usage rules for Go

## Supported abstractions

Use the supported types as follows:

| Type | Use for | Avoid when |
| --- | --- | --- |
| `Option[T]` | A value may be absent and absence is not an error. | Zero value already has the desired meaning. |
| `Result[T]` | A computation may fail with an `error`. | Both alternatives are valid non-error domain states. |
| `Either[L, R]` | A value is exactly one of two valid alternatives. | One side is just `error`; use `Result[T]`. |
| `Either3[T1, T2, T3]` | A value is exactly one of three valid alternatives. | More than five variants are needed; prefer an interface or domain type. |
| `Either4[T1, T2, T3, T4]` | A value is exactly one of four valid alternatives. | More than five variants are needed; prefer an interface or domain type. |
| `Either5[T1, T2, T3, T4, T5]` | A value is exactly one of five valid alternatives. | More than five variants are needed; prefer an interface or domain type. |
| `Future[T]` | An async value that will resolve or reject. | `context.Context`, goroutines, channels, or `errgroup` express the workflow better. |
| `IO[T]` | A lazy synchronous side-effecting computation. | The side effect should run immediately and directly. |
| `IOEither[T]` | A lazy synchronous side-effecting computation that can fail. | Plain `(T, error)` is clearer at the call site. |
| `Task[T]` | A lazy async computation. | The lifecycle needs explicit cancellation, fan-out, or structured concurrency. |
| `TaskEither[T]` | A lazy async computation that can fail. | Standard Go concurrency primitives are clearer. |
| `State[S, A]` | A stateful computation with explicit state threading. | A normal mutable local variable is clearer. |

## General decision rules

1. Prefer `Option[T]` over `*T` when absence is a meaningful part of the API and callers must handle it explicitly.
2. Prefer `Result[T]` inside composition-heavy domain logic. At public Go API boundaries, accept and return idiomatic `(T, error)` unless the project already standardizes on `Result[T]`.
3. Prefer `Either[L, R]` and `Either3` through `Either5` only when every variant is a valid domain outcome. If one variant is failure, model it with `Result[T]`.
4. Prefer direct methods for same-type transformations. Use sub-packages for transformations that change type parameters, because Go methods cannot introduce their own type parameters.
5. Do not use `MustGet` as normal control flow. Use `Get`, `OrElse`, `OrEmpty`, `Match`, `Map`, or `FlatMap`. `MustGet` is acceptable only when the value is guaranteed by construction, in tests, or inside a deliberate `mo.Do` block.
6. Keep conversions at boundaries. Convert `map` lookups, pointer values, nullable database fields, JSON nullable fields, or `(T, error)` returns into monadic types near the boundary; unwrap before calling code that expects idiomatic Go values.

## Option rules

1. Construct options with:
```go
some := mo.Some("alice")
none := mo.None[string]()
fromPtr := mo.PointerToOption(ptr)
fromMap := mo.TupleToOption(m[key])
fromEmptyable := mo.EmptyableToOption(value)
```
2. Extract safely:
```go
value, ok := some.Get()
value = none.OrElse("anonymous")
value = none.OrEmpty()
```
3. Use `Map` or `FlatMap` for same-type chains:
```go
normalized := mo.Some("alice").Map(func(s string) (string, bool) {
    return strings.ToUpper(s), true
})
```
4. Use `option.Map`, `option.FlatMap`, and `option.Pipe1` through `option.Pipe10` when a chain changes types:
```go
import "github.com/samber/mo/option"
out := option.Pipe2(
    mo.Some(42),
    option.Map(func(v int) string { return strconv.Itoa(v) }),
    option.Map(func(s string) []byte { return []byte(s) }),
)
```
5. Use `IsSome` or `IsNone` in new code when naming should mirror Rust-style `Option`. `IsPresent` and `IsAbsent` are also available and equivalent aliases in practice.
6. `Option[T]` supports JSON, text, binary, gob, SQL scanning, and SQL driver value interfaces. Use it directly in structs only when the persistence or API contract accepts nullable fields.
7. MUST NOT use `Option[*T]` or `Option[error]`. Use `*T` or `error` directly instead.

## Result rules

1. Construct results with:
```go
ok := mo.Ok(42)
errResult := mo.Err[int](err)
formatted := mo.Errf[int]("invalid id: %s", id)
fromCall := mo.TupleToResult(os.ReadFile(path))
tried := mo.Try(func() ([]byte, error) { return os.ReadFile(path) })
```
2. Extract safely:
```go
value, err := fromCall.Get()
value = fromCall.OrElse(defaultValue)
```
3. Use `Result[T]` to compose fallible operations of the same value type:
```go
result := mo.Ok(21).Map(func(v int) (int, error) {
    return v * 2, nil
})
```
4. Use `result.Map`, `result.FlatMap`, and `result.Pipe1` through `result.Pipe10` when the pipeline changes types:
```go
import "github.com/samber/mo/result"
parsed := result.Pipe2(
    mo.TupleToResult(os.ReadFile("config.json")),
    result.Map(func(data []byte) Config { return parseConfig(data) }),
    result.FlatMap(func(cfg Config) mo.Result[Config] { return validateConfig(cfg) }),
)
```
5. Use `MapErr` to recover or transform failures. Do not swallow errors silently; preserve or wrap context where it helps diagnosis.
6. Use `ToEither` only when a downstream API expects `Either[error, T]`.
7. Use `mo.Do` sparingly. It captures panics produced by `MustGet` and returns a `Result[T]`, but it can hide ordinary Go control flow if overused.

## Either rules

1. Construct two-way alternatives with:
```go
left := mo.Left[CachedUser, FreshUser](cached)
right := mo.Right[CachedUser, FreshUser](fresh)
```
2. Use `Left`, `Right`, `LeftOrElse`, `RightOrElse`, `Match`, `Swap`, and `Unpack` to consume values explicitly.
3. Use `either.MapLeft`, `either.MapRight`, `either.FlatMapLeft`, `either.FlatMapRight`, `either.Match`, and `either.Pipe1` through `either.Pipe10` for type-changing pipelines.
4. For `3`, `4`, or `5` alternatives, use `Either3`, `Either4`, or `Either5` constructors and the matching `either3`, `either4`, or `either5` sub-package. Do not encode complex business state as `Either5` if a named struct, enum-like type, or interface would be more readable.

## Future, IO, Task, and State rules

1. Use `Future[T]` for promise-like async values. Always call `Collect`, `Result`, or `Either` at a clear boundary. Use `Catch`, `Then`, and `Finally` for linear async chains, not for complex orchestration.
2. Use `IO[T]` and `IOEither[T]` to delay side effects until `Run`. This is useful in tests or dependency injection seams; it is unnecessary for straightforward code.
3. Use `Task[T]` and `TaskEither[T]` for lazy async computations. Keep cancellation and ownership explicit in surrounding code.
4. Use `State[S, A]` only when state threading is the central abstraction. Otherwise, a simple local variable is usually more idiomatic.

## Pipeline rules

1. The pipeline sub-packages are github.com/samber/mo: option, result, either, either3, either4, either5.
2. Use `Pipe1` through `Pipe10` for readable left-to-right transformations. Do not nest `Map` calls when a pipe makes the data flow clearer.
3. Choose the smallest pipeline that fits the number of steps. Avoid forcing a long pipeline when named intermediate variables make debugging easier.

## Review checklist

1. `Option[T]` is used only for meaningful absence, not as a replacement for every pointer.
2. `Result[T]` does not leak into public APIs unless the project convention explicitly accepts it.
3. `Either` is not used as a disguised error channel.
4. No `MustGet` used in production code, because it can panic.
5. Type-changing transformations use the appropriate sub-package instead of impossible direct method chaining.
6. Monadic code reduces nil/error branching rather than making simple Go code harder to read.