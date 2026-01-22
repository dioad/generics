package generics

import "slices"

// Filter returns a new slice containing only the elements that satisfy the predicate.
// If the input slice is empty or nil, it returns the input slice.
func Filter[A any](arr []A, predicate func(A) bool) []A {
	if len(arr) == 0 {
		return arr
	}

	result := make([]A, 0, len(arr))

	for _, a := range arr {
		if predicate(a) {
			result = append(result, a)
		}
	}

	return result
}

// Reduce applies a function to each element in a slice, accumulating a single result.
// It starts with the initial value and sequentially applies f to the accumulator and each element.
func Reduce[A any, B any](arr []A, initial B, f func(B, A) B) B {
	result := initial

	for _, a := range arr {
		result = f(result, a)
	}

	return result
}

// Contains returns true if the slice contains an element that satisfies the predicate.
func Contains[T any](arr []T, f func(T) bool) bool {
	return slices.IndexFunc(arr, f) != -1
}

// ForEach applies a function to each element of a slice and returns the first error encountered.
//
// Unlike Apply, ForEach stops processing as soon as an error is encountered.
func ForEach[A any](arr []A, f func(A) error) error {
	for _, a := range arr {
		if err := f(a); err != nil {
			return err
		}
	}
	return nil
}
