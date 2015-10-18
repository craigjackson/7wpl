package models

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite3", "file:database.sqlite3")
	if err != nil {
		log.Fatalf("sql.Open: %s\n", err.Error())
	}
}
