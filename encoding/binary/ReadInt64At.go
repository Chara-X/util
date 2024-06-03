package binary

import (
	"encoding/binary"
	"io"
)

func ReadInt64At(r io.ReaderAt, off int64) int64 {
	var buf = make([]byte, 8)
	r.ReadAt(buf, off)
	return int64(binary.LittleEndian.Uint64(buf))
}
