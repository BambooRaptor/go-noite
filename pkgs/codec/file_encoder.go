package codec

import (
	"fmt"
	"os"

	"github.com/BambooRaptor/go-noite/pkgs/database"
)

type FileEncoder struct {
	file *os.File
}

type FileDecoder struct {
	file *os.File
}

func NewFileEncoder(file *os.File) DatabaseEncoder {
	return &FileEncoder{file}
}

func (fe *FileEncoder) Encode(db *database.Database) error {
	header := fmt.Sprintf("noite %v all your base are belong to us.\n", db.Version())
	var headerBytes [64]byte
	copy(headerBytes[:], header)
	_, err := fe.file.WriteAt(headerBytes[:], 0)
	return err
}

func NewFileDecoder(file *os.File) DatabaseDecoder {
	return &FileDecoder{file}
}

func (fd *FileDecoder) Decode() (*database.Database, error) {

	// Read Header Info => Stored in a 256 chunk at the start
	buffer := make([]byte, 256)
	read, err := fd.file.Read(buffer)
	if err != nil {
		return nil, err
	}
	fmt.Printf("[DATABASE HEADER] %v\n", buffer[0:read])
	return database.NewDatabase(), nil
}
