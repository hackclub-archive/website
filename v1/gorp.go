package v1

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/coopernurse/gorp"
	"github.com/hackedu/backend/v1/model"
	_ "github.com/lib/pq"
)

var Dbm *gorp.DbMap

func init() {
	Dbm = newDbMap()

	Dbm.AddTable(model.School{}).SetKeys(true, "Id")
	Dbm.AddTableWithName(model.User{}, "Users").SetKeys(true, "Id")
	Dbm.AddTable(model.Application{}).SetKeys(true, "Id")
}

func newDbMap() *gorp.DbMap {
	dialect, driver := dialectAndDriver()
	return &gorp.DbMap{Db: connect(driver), Dialect: dialect}
}

func dialectAndDriver() (gorp.Dialect, string) {
	return gorp.PostgresDialect{}, "postgres"
}

func connect(driver string) *sql.DB {
	dsn := fmt.Sprintf("postgres://docker:docker@%s/docker",
		os.Getenv("DB_1_PORT_5432_TCP_ADDR"))

	if os.Getenv("ENV") == "PRODUCTION" {
		dsn = os.Getenv("DB_DSN")
		if dsn == "" {
			panic("DB_DSN env variable not set")
		}
	}

	db, err := sql.Open(driver, dsn)
	if err != nil {
		panic("Error connecting to db: " + err.Error())
	}
	return db
}
