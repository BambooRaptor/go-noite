package v0

import (
	"io"

	"github.com/BambooRaptor/go-noite/pkgs/database"
)

type Writer struct {
	version     Ui16
	header      String
	collections List[collection]
}

func NewWriter(db *database.Database) *Writer {
	return &Writer{
		version: new(Ui16).New(7),
		header:  MakeString("all your noite are belong to us."),
		collections: new(List[collection]).New(
			collection{
				new(Ui64).New(5664),
				MakeString("hello"),
			},
			collection{
				new(Ui64).New(4554),
				MakeString("world"),
			},
		),
	}
}

func (s Writer) GetSchema() []io.WriterTo {
	return compound(s.version, s.header, s.collections)
}

type collection struct {
	dataOffset Ui64
	name       String
}

func (col collection) GetSchema() []io.WriterTo {
	return compound(col.dataOffset, col.name)
}
