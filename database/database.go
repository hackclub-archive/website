package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Postgres driver
)

var db *sqlx.DB

// Init initializes the internal database handle.
func Init(name, datasource string) error {
	var err error
	db, err = sqlx.Open(name, datasource)
	if err != nil {
		return err
	}

	return nil
}

// Close database connection
func Close() {
	db.Close()
}
