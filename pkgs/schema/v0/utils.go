package v0

import (
	"encoding/binary"
	"io"
	"slices"
	"unsafe"

	"github.com/BambooRaptor/go-noite/pkgs/schema"
)

func compound(blocks ...schema.SchemaObject) []io.WriterTo {
	slcs := make([][]io.WriterTo, 0)
	for _, blk := range blocks {
		slcs = append(slcs, blk.GetSchema())
	}
	return slices.Concat(slcs...)
}

func writeBinaryTo(w io.Writer, data any) (n int64, err error) {
	err = binary.Write(w, binary.LittleEndian, data)
	if err == nil {
		n = int64(unsafe.Sizeof(data))
	}
	return n, err
}

func writers(writers ...io.WriterTo) []io.WriterTo {
	return writers
}
