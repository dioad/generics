// Package generics provides generic utility functions for working with Go types.
//
// This package implements common functional programming patterns like Map, Filter,
// and Reduce using Go's generics support. It aims to provide a simple, consistent
// API for working with collections in Go.
package generics

import (
	"errors"
	"fmt"
)

var (
	// ErrDifferentLength is returned when two slices of different lengths are provided
	// to a function that requires equal lengths.
	ErrDifferentLength = errors.New("arrays must be of equal length")

	// ErrNotFound is returned when an element matching a predicate is not found.
	ErrNotFound = errors.New("not found")
)

// SafeApply applies a function to each element of a slice.
// This is useful for side effects like logging or metrics collection.
func SafeApply[A any](f func(A), arr []A) {
	for _, a := range arr {
		f(a)
	}
}

// Apply applies a function to each element of a slice and collects all errors.
// It returns a MapError containing all non-nil errors returned by f, keyed by index.
// If all calls to f return nil, it returns nil.
func Apply[A any](f func(A) error, arr []A) error {
	err := NewMapError()

	for i, a := range arr {
		e := f(a)
		if e != nil {
			err.Add(i, e)
		}
	}

	if err.HasError() {
		return err
	}

	return nil
}

// MapError is a collection of errors returned by functions applied to slices.
// It maps the index of the element that caused the error to the error itself.
type MapError struct {
	Errors map[int]error
}

// Error returns a string representation of the MapError.
func (m *MapError) Error() string {
	return fmt.Sprintf("%d errors", len(m.Errors))
}

// HasError returns true if the MapError contains any errors.
func (m *MapError) HasError() bool {
	return len(m.Errors) > 0
}

// Add adds an error to the MapError at the specified index.
func (m *MapError) Add(idx int, err error) {
	m.Errors[idx] = err
}

// NewMapError creates a new, empty MapError.
func NewMapError() *MapError {
	return &MapError{
		Errors: make(map[int]error),
	}
}

// SafeMap applies a function to each element of a slice and returns a slice of the results.
func SafeMap[A any, B any](f func(A) B, arr []A) []B {
	results := make([]B, len(arr))

	for i, a := range arr {
		results[i] = f(a)
	}

	return results
}

// Map applies a function that can return an error to each element of a slice.
// It returns a slice of all results and a MapError if any call to f returned an error.
func Map[A any, B any](f func(A) (B, error), arr []A) ([]B, error) {
	results := make([]B, len(arr))

	err := NewMapError()

	for i, a := range arr {
		b, e := f(a)
		if e != nil {
			err.Add(i, e)
		}
		results[i] = b
	}

	if err.HasError() {
		return results, err
	}

	return results, nil
}

// Pair is a simple tuple of two values of possibly different types.
type Pair[A, B any] struct {
	A A
	B B
}

// Compact returns a new slice with all zero values removed.
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

// Zip combines two slices into a single slice of Pairs.
// It returns ErrDifferentLength if the slices are not the same length.
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

// SelectOne returns the first element in a slice that satisfies the predicate.
// If no such element is found, it returns the zero value and ErrNotFound.
func SelectOne[T any](arr []T, f func(T) bool) (T, error) {
	var zero T
	for _, v := range arr {
		if f(v) {
			return v, nil
		}
	}

	return zero, ErrNotFound
}
