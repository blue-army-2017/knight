package controller

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {
	dbName, dbNameExists := os.LookupEnv("DB_NAME")
	if !dbNameExists {
		dbName = "knight.db"
	}

	var err error
	db, err = sql.Open("sqlite3", dbName)
	if err != nil {
		panic(err)
	}
}
