package codec

import (
	"encoding/binary"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/BambooRaptor/go-noite/pkgs/database"
	"github.com/BambooRaptor/go-noite/pkgs/schema/v0"
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
	writer := v0.NewWriter(db)
	for _, w := range writer.GetSchema() {
		_, err := w.WriteTo(fe.file)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewFileDecoder(file *os.File) DatabaseDecoder {
	return &FileDecoder{file}
}

const header_size = 48

type header [header_size]byte

func (fd *FileDecoder) Decode() (*database.Database, error) {
	// Read Header Info => Stored in a 64 chunk at the start
	var headerBytes header
	err := binary.Read(fd.file, binary.LittleEndian, &headerBytes)
	if err != nil {
		fmt.Printf("[BINARY READ ERROR] %q\n", err)
		return nil, err
	}
	fmt.Printf("[DATABASE HEADER] %q\n", headerBytes[:])
	version, err := getVersionFromHeader(headerBytes)
	fmt.Printf("[VERSION] %v\n", version)

	if err != nil {
		fmt.Printf("[HEADER ERROR] Error parsing header %q\n", headerBytes)
		return nil, err
	}

	return database.NewDatabase(), nil
}

func parseEncodingSchema(str string) (uint64, error) {
	return strconv.ParseUint(str, 10, 1)
}

func getVersionFromHeader(header [header_size]byte) (uint64, error) {
	versionStr := strings.Split(string(header[:]), " ")[1]
	return parseEncodingSchema(versionStr)
}
