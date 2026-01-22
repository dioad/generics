package generics

import "reflect"

// IsZeroValue returns true if the value v is the zero value of its type.
// It handles nil interfaces, pointers, slices, and maps, as well as primitive types and structs.
func IsZeroValue[T any](v T) bool {
	t := reflect.TypeOf(v)
	// Handle nil interfaces
	if t == nil {
		return true
	}
	return reflect.DeepEqual(v, reflect.Zero(t).Interface())
}
