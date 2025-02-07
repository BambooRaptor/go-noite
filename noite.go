package noite

import (
	"fmt"
	"os"
)

// Let's think through this.
// What would the header information for the database include?
// "noite v0.1 all your base are belong to us"
type Database struct{}

func NewDatabase() *Database {
	return new(Database)
}

func CreateOrOpen(path string, file string) (*Database, error) {
	fl, err := os.Open(path + file)
	if err != nil {
		fmt.Printf("[FILE OPEN ERROR] Error opening file %q.\n> %v\n", path+file, err)
		fmt.Printf("[==>] Attempting to create file %q\n", path)

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

	}
	defer fl.Close()
	return NewDatabase(), nil
}
