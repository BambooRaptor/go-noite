package database

import "fmt"

// Let's think through this.
// What would the header information for the database include?
// "noite v0.1 all your base are belong to us"

type Version struct {
	Major uint8
	Minor uint8
	Patch uint8
}

func LatestVersion() Version {
	return Version{0, 0, 0}
}

type info struct {
	version Version
}

type Database struct {
	info info
}

type DatabaseConfig struct {
	Version Version
}

func NewVersion(maj uint8, min uint8, pat uint8) Version {
	return Version{maj, min, pat}
}

func LatestConfig() DatabaseConfig {
	return DatabaseConfig{LatestVersion()}
}

func ConfigWithVersion(version Version) DatabaseConfig {
	return DatabaseConfig{version}
}

func NewDatabase(cfg DatabaseConfig) *Database {
	return &Database{
		info{cfg.Version},
	}
}

func (db *Database) Version() string {
	return fmt.Sprintf("v%d.%d.%d", db.info.version.Major, db.info.version.Minor, db.info.version.Patch)
}
