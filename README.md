# Generics

A collection of generic utility functions for Go (1.22+). This package provides common functional programming patterns
and slice utilities using Go's generics support.

## Installation

```bash
go get github.com/dioad/generics
```

## Features

- **Functional Patterns**: `Map`, `Filter`, `Reduce`, `ForEach`.
- **Slice Utilities**: `Compact`, `Zip`, `SelectOne`.
- **Type Utilities**: `IsZeroValue`.
- **Error Handling**: `MapError` for collecting multiple errors during batch operations.

## Usage

### Map

Apply a function to each element of a slice and return a new slice of results.

```go
import "github.com/dioad/generics"

nums := []int{1, 2, 3}
doubled := generics.SafeMap(func(x int) int {
    return x * 2
}, nums)
// [2, 4, 6]
```

### Filter

Return a new slice containing only elements that satisfy a predicate.

```go
nums := []int{1, 2, 3, 4}
evens := generics.Filter(nums, func(x int) bool {
    return x % 2 == 0
})
// [2, 4]
```

### Reduce

Accumulate a single result from a slice.

```go
nums := []int{1, 2, 3, 4}
sum := generics.Reduce(nums, 0, func(acc, x int) int {
    return acc + x
})
// 10
```

### IsZeroValue

Check if a value is the zero value of its type.

```go
var s string
generics.IsZeroValue(s) // true

s = "hello"
generics.IsZeroValue(s) // false
```

## Documentation

Full documentation is available on [pkg.go.dev](https://pkg.go.dev/github.com/dioad/generics).

## License

Apache License 2.0
