package unsafe

import "unsafe"

func StringData(str string) *byte {
	return (*struct {
		ptr *byte
		len int
	})(unsafe.Pointer(&str)).ptr
}
