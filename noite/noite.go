package noite

import (
	"fmt"
	"os"

	"github.com/BambooRaptor/go-noite/pkgs/codec"
	"github.com/BambooRaptor/go-noite/pkgs/database"
)

func CreateOrOpen(path string, file string) (*database.Database, error) {
	// Try opening the file
	fl, err := os.Open(path + file)

	fmt.Printf("[OPENING FILE]\n")
	// If the file doesn't exist
	if err != nil {
		fmt.Printf("> %v\n", err)
		fmt.Printf("[==>] Attempting to create file %q\n", path+file)

		// Create all folders
		err = os.MkdirAll(path, 0700)
		if err != nil {
			fmt.Printf("[PATH CREATE ERROR] Error creating folders %q.\n", path)
			return nil, err
		}

		// Finally, create the file
		fl, err = os.Create(path + file)
		if err != nil {
			fmt.Printf("[FILE CREATE ERROR] Error creating file %q.\n", path+file)
			return nil, err
		}
		defer fl.Close()
		fmt.Printf("[<==] Successfully created file %q\n", path+file)

		db := database.NewDatabase(database.LatestConfig())

		// Return the database and encode it in the file
		return db, encodeDatabase(fl, db)
	}
	defer fl.Close()

	return decodeDatabase(fl)
}

func encodeDatabase(file *os.File, db *database.Database) error {
	encoder := codec.NewFileEncoder(file)
	return encoder.Encode(db)
}

func decodeDatabase(file *os.File) (*database.Database, error) {
	decoder := codec.NewFileDecoder(file)
	return decoder.Decode()
}
