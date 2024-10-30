package unsafe

import "unsafe"

func Slice[T any](ptr *T, len int) []T {
	return *(*[]T)(unsafe.Pointer(&struct {
		ptr *T
		len int
		cap int
	}{ptr, len, len}))
}
