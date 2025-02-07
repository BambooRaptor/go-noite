package database

import "fmt"

// Let's think through this.
// What would the header information for the database include?
// "noite v0.1 all your base are belong to us"

type version struct {
	major uint
	minor uint
}

type info struct {
	version version
}

type Database struct {
	info info
}

func NewDatabase() *Database {
	return &Database{
		info{
			version{0, 1},
		},
	}
}

func (db *Database) Version() string {
	return fmt.Sprintf("v%d.%d", db.info.version.major, db.info.version.minor)
}
