package generics

import (
	"errors"
	"fmt"
)

var (
	ErrDifferentLength = errors.New("arrays must be of equal length")
)

// Apply applies a function to each element of an array
//
// # It returns an array of errors, where each error corresponds to the result of the function applied to the element at the same index
//
// It returns nil if all function calls return nil
func Apply[A any](f func(A) error, arr []A) error {
	err := NewMapError[A]()

	for _, a := range arr {
		e := f(a)
		if e != nil {
			err.Add(a, e)
		}
	}

	if err.HasError() {
		return err
	}

	return nil
}

// MapError is a collection of errors
//
// It is used to collect errors from a function applied to each element of an array
type MapError[T any] struct {
	Errors []Pair[T, error]
}

// Error returns a string representation of the MapError
func (m *MapError[T]) Error() string {
	return fmt.Sprintf("%d errors", len(m.Errors))
}

// HasError returns true if the MapError has errors
func (m *MapError[T]) HasError() bool {
	return len(m.Errors) > 0
}

// Add adds an error to the MapError
func (m *MapError[T]) Add(a T, err error) {
	m.Errors = append(m.Errors, Pair[T, error]{
		A: a,
		B: err,
	})
}

// NewMapError creates a new MapError
func NewMapError[T any]() *MapError[T] {
	return &MapError[T]{
		Errors: make([]Pair[T, error], 0),
	}
}

// Map applies a function to each element of an array
//
// It returns an array of results and an array of errors, where each result and error
// corresponds to the result of the function applied to the element at the same index
//
// It returns nil if all function calls return nil
func Map[A any, B any](f func(A) (B, error), arr []A) ([]B, error) {
	results := make([]B, len(arr))

	err := NewMapError[A]()

	for i, a := range arr {
		b, e := f(a)
		if e != nil {
			err.Add(a, e)
		}
		results[i] = b
	}

	if err.HasError() {
		return results, err
	}

	return results, nil
}

// Pair is a tuple of two values
type Pair[A, B any] struct {
	A A
	B B
}

// Compact removes zero values from an array
func Compact[A comparable](arr []A) []A {
	var zero A

	result := make([]A, 0)

	for _, a := range arr {
		if a != zero {
			result = append(result, a)
		}
	}

	return result
}

func Zip[A any, B any](a []A, b []B) ([]Pair[A, B], error) {
	if len(a) != len(b) {
		return nil, ErrDifferentLength
	}

	result := make([]Pair[A, B], len(a))

	for i := 0; i < len(a); i++ {
		pair := Pair[A, B]{
			A: a[i],
			B: b[i],
		}
		result[i] = pair
	}

	return result, nil
}
