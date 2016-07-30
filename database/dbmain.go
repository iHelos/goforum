package database

import (
	//"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "user=ihelos dbname=goforum password=P1rate sslmode=disable")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

func GetDB() *sql.DB {
	return db
}