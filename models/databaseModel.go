package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var database *sql.DB

//DatabaseConnect returns database connection
func DatabaseConnect() {
	conStr := "host=95.85.122.37 port=5432 user=mobia password=mobia dbname=mobia sslmode=disable"
	db, err := sql.Open("postgres", conStr)
	if err != nil {
		panic(err)
	}
	database = db
}

//GetDB returns database object
func GetDB() *sql.DB {
	return database
}
