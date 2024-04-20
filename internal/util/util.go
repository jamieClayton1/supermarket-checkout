package util

// Generate a pointer for the given value
func Pointer[T interface{}](val T) *T {
	return &val
}
