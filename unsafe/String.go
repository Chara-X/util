package unsafe

import "unsafe"

func String(ptr *byte, len int) string {
	return *(*string)(unsafe.Pointer(&struct {
		ptr *byte
		len int
	}{ptr, len}))
}
