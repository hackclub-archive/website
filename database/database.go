package database

import (
	"database/sql"
	_ "github.com/lib/pq" // Postgres driver
)

var db *sql.DB

// Init initializes the internal database handle.
func Init(name, datasource string) error {
	var err error
	db, err = sql.Open(name, datasource)
	if err != nil {
		return err
	}

	return nil
}

// Close database connection
func Close() {
	db.Close()
}
