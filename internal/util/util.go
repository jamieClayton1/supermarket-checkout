package util

func Pointer[T interface{}](val T) *T {
	return &val
}
