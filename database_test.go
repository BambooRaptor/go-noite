package noite_test

import (
	"errors"
	"os"
	"testing"

	"github.com/BambooRaptor/go-noite"
)

func TestDatabaseCreation(t *testing.T) {
	path := "./tmp/db/"
	file := "test.noite"
	_, err := noite.CreateOrOpen(path, file)
	if err != nil {
		t.Fatalf("Error opening Database %q.\n> %v\n", path, err)
	}

	// Database has either been created or opened. Check if the file exists
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		t.Fatalf("Database with path %q not created.\n> %v\n", path, err)
	}

	// Cleanup
	if err = os.Remove(path + file); err != nil {
		t.Fatalf("Error removing tmp database %q.\n> %v\n", path, err)
	}
}
