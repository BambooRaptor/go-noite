package noite_test

import (
	"os"
	"testing"

	"github.com/BambooRaptor/go-noite/noite"
	"github.com/BambooRaptor/go-noite/pkgs/database"
)

const path = "../tmp/db/"
const file = "test.noite"

func cleanTempFile(path string, file string) {
	_ = os.Remove(path + file)
}

func cleanTemp() {
	cleanTempFile(path, file)
}

func TestDatabaseFileCreation(t *testing.T) {
	cleanTemp()

	// Should create a new database file
	t.Run("create new database file", func(t *testing.T) {
		createOrOpen(t, path, file)
	})

	// Should open the created database file
	t.Run("open existing database file", func(t *testing.T) {
		createOrOpen(t, path, file)
	})
}

func createOrOpen(t *testing.T, path string, file string) *database.Database {
	db, err := noite.CreateOrOpen(path, file)
	if err != nil {
		t.Fatalf("[TEST ERROR] Error creating Database %q.\n> %v\n", path, err)
	}
	return db
}

func TestDatabaseCollections(t *testing.T) {
	cleanTemp()
	db := createOrOpen(t, path, file)
	db.Collection("hello")
}
