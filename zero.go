package generics

import "reflect"

// IsZeroValue returns true if the value is the zero value of its type.
func IsZeroValue[T any](a T) bool {
	var zero T
	return reflect.DeepEqual(a, zero)
}
