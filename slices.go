package generics

import (
	"errors"
)

var (
	ErrDifferentLength = errors.New("arrays must be the same length")
)

// Apply applies a function to each element of an array
//
// # It returns an array of errors, where each error corresponds to the result of the function applied to the element at the same index
//
// It returns nil if all function calls return nil
func Apply[A any](f func(A) error, arr []A) []error {
	errors := make([]error, len(arr))

	errFound := false

	for i, a := range arr {
		err := f(a)
		if err != nil {
			errFound = true
			errors[i] = err
		}
	}

	if errFound {
		return errors
	}

	return nil
}

// Map applies a function to each element of an array
//
// It returns an array of results and an array of errors, where each result and error
// corresponds to the result of the function applied to the element at the same index
//
// It returns nil if all function calls return nil
func Map[A any, B any](f func(A) (B, error), arr []A) ([]B, []error) {
	results := make([]B, len(arr))

	errors := make([]error, len(arr))

	errFound := false
	for i, a := range arr {
		b, e := f(a)
		if e != nil {
			errFound = true
			errors[i] = e
		}
		results[i] = b
	}

	if errFound {
		return results, errors
	}

	return results, nil
}

type Pair[A, B any] struct {
	A A
	B B
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
