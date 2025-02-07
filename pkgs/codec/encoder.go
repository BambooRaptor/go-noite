package codec

import (
	"fmt"

	"github.com/BambooRaptor/go-noite/pkgs/database"
)

const LATEST_ENCODING_SCHEMA uint64 = 0

type DatabaseEncoder interface {
	Encode(*database.Database) error
}

type DatabaseDecoder interface {
	Decode() (*database.Database, error)
}

type EncodingSchema uint16

func (es EncodingSchema) Version() string {
	return fmt.Sprintf("v%d", es)
}
