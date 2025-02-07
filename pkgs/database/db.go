package database

// Let's think through this.
// What would the header information for the database include?
// "noite v0.1 all your base are belong to us"

type Database struct {
	collections map[string]*Collection
}

func NewDatabase() *Database {
	return &Database{make(map[string]*Collection)}
}

type Collection struct {
	docs map[string]map[string]string
}

func newCollection() *Collection {
	return &Collection{make(map[string]map[string]string)}
}

func (db *Database) Collection(col string) *Collection {
	collection, exists := db.collections[col]
	if !exists {
		collection = newCollection()
		db.collections[col] = collection
	}
	return collection
}
