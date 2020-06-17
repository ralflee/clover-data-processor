package db

import (
	"database/sql"
	"log"

	//_ "github.com/lib/pq"
	_ "github.com/proullon/ramsql/driver"
)

var db *sql.DB

//InitDB initalize database connection
func InitDB(dataSourceName string) (*sql.DB, error) {
	//db, err := sql.Open("postgres", dataSourceName)
	db, err := sql.Open("ramsql", "TestLoadUserAddresses")
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	return db, nil
}
