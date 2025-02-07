package v0

import (
	"fmt"
	"io"
	"math"
	"slices"

	"github.com/BambooRaptor/go-noite/pkgs/schema"
)

type String struct {
	length Ui16
	text   ByteSlice
}

func MakeString(str string) String {
	strlen := len(str)
	max := math.MaxUint16
	if strlen >= max {
		panic(fmt.Sprintf("Cannot have a String be more than %d characters", max))
	}

	return String{
		length: new(Ui16).New(uint16(strlen)),
		text:   new(ByteSlice).New([]byte(str)),
	}
}

func (s String) GetSchema() []io.WriterTo {
	return slices.Concat(s.length.GetSchema(), s.text.GetSchema())
}

type List[T schema.SchemaObject] struct {
	total Ui16
	list  []T
}

func (lo List[T]) GetSchema() []io.WriterTo {
	list := make([][]io.WriterTo, 0)
	for _, so := range lo.list {
		list = append(list, so.GetSchema())
	}
	all := slices.Concat(list...)
	return slices.Concat(lo.total.GetSchema(), all)
}

func (lo *List[T]) New(els ...T) List[T] {
	lo.list = els
	return *lo
}
