package database

import (
	//"fmt"
	"database/sql"
	_ "github.com/bmizerany/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "user=ihelos dbname=goforum sslmode=disable")
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