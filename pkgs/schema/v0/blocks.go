package v0

import (
	"io"
)

type SchemaWriter[T any] struct {
	val T
}

func (sw *SchemaWriter[T]) New(val T) SchemaWriter[T] {
	sw.val = val
	return *sw
}

func (sw SchemaWriter[T]) WriteTo(w io.Writer) (n int64, err error) {
	return writeBinaryTo(w, sw.val)
}

func (sw SchemaWriter[T]) GetSchema() []io.WriterTo {
	return writers(sw)
}

type Ui16 = SchemaWriter[uint16]
type Ui64 = SchemaWriter[uint64]

// used for strings
type ByteSlice = SchemaWriter[[]byte]
