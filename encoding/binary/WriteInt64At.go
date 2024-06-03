package binary

import (
	"encoding/binary"
	"io"
)

func WriteInt64At(w io.WriterAt, val int64, off int64) {
	var buf = make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(val))
	w.WriteAt(buf, off)
}
