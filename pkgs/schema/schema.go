package schema

import "io"

type SchemaObject interface {
	GetSchema() []io.WriterTo
}
