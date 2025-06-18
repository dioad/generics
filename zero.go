package generics

import "reflect"

// IsZeroValue returns true if the value is the zero value of its type.
func IsZeroValue[T any](a T) bool {
	t := reflect.TypeOf(a)
	// Handle nil interfaces
	if t == nil {
		return true
	}
	return reflect.DeepEqual(a, reflect.Zero(t).Interface())
}
