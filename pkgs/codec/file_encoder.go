package codec

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/BambooRaptor/go-noite/pkgs/database"
)

type FileEncoder struct {
	file *os.File
}

type FileDecoder struct {
	file *os.File
}

const HEADER_BUFFER = 64

func NewFileEncoder(file *os.File) DatabaseEncoder {
	return &FileEncoder{file}
}

func (fe *FileEncoder) Encode(db *database.Database) error {
	header := fmt.Sprintf("noite %v all your base are belong to us.\n", db.Version())
	var headerBytes [HEADER_BUFFER]byte
	copy(headerBytes[:], header)
	_, err := fe.file.WriteAt(headerBytes[:], 0)
	return err
}

func NewFileDecoder(file *os.File) DatabaseDecoder {
	return &FileDecoder{file}
}

func (fd *FileDecoder) Decode() (*database.Database, error) {

	// Read Header Info => Stored in a 256 chunk at the start
	header := [HEADER_BUFFER]byte{}
	read, err := fd.file.Read(header[:])
	if err != nil {
		return nil, err
	}
	fmt.Printf("[DATABASE HEADER] %q\n", header[0:read])
	version, err := getVersionFromHeader(header)

	if err != nil {
		fmt.Printf("[HEADER ERROR] Error parsing header %q\n", header)
		return nil, err
	}

	return database.NewDatabase(database.ConfigWithVersion(*version)), nil
}

func parseVersion(str string) (*database.Version, error) {
	if str[0] == 'v' {
		str = str[1:]
		versions := strings.Split(str, ".")
		if len(versions) == 3 {
			major, majErr := strconv.ParseUint(versions[0], 10, 8)
			minor, minErr := strconv.ParseUint(versions[1], 10, 8)
			patch, patErr := strconv.ParseUint(versions[2], 10, 8)
			if majErr == nil && minErr == nil && patErr == nil {
				ver := database.NewVersion(uint8(major), uint8(minor), uint8(patch))
				return &ver, nil
			}
		}
	}

	return nil, fmt.Errorf("%q is not a valid version string", str)
}

func getVersionFromHeader(header [HEADER_BUFFER]byte) (*database.Version, error) {
	versionStr := strings.Split(string(header[:]), " ")[1]
	return parseVersion(versionStr)
}
