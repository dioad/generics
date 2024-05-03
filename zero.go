package generics

import "reflect"

func IsZeroValue[T any](a T) bool {
	var zero T
	return reflect.DeepEqual(a, zero)
}
