package codec

import "github.com/BambooRaptor/go-noite/pkgs/database"

type DatabaseEncoder interface {
	Encode(*database.Database) error
}

type DatabaseDecoder interface {
	Decode() (*database.Database, error)
}
