package noite_test

import (
	"errors"
	"io/fs"
	"os"
	"testing"

	"github.com/BambooRaptor/go-noite/noite"
)

func TestDatabaseFileCreation(t *testing.T) {
	path := "../tmp/db/"
	file := "test.noite"

	// Clean the tmp file for testing
	if err := os.Remove(path + file); !errors.Is(err, fs.ErrNotExist) {
		t.Fatalf("Error removing tmp database %q.\n> %v\n", path, err)
	}

	// Should create
	t.Run("create new database file", func(t *testing.T) {
		_, err := noite.CreateOrOpen(path, file)
		if err != nil {
			t.Fatalf("[TEST ERROR] Error creating Database %q.\n> %v\n", path, err)
		}
	})

	t.Run("open existing database file", func(t *testing.T) {
		_, err := noite.CreateOrOpen(path, file)
		if err != nil {
			t.Fatalf("[TEST ERROR] Error opening Database %q.\n> %v\n", path, err)
		}
	})
}
