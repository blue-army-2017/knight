package controller

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {
	dbName, dbNameExists := os.LookupEnv("DB_NAME")
	if !dbNameExists {
		dbName = "knight.db"
	}
	log.Printf("Connecting to database %s", dbName)

	var err error
	db, err = sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Successfully connected to database %s", dbName)

	m, err := migrate.New("file://schema", fmt.Sprintf("sqlite3://%s", dbName))
	if err != nil {
		log.Fatalln(err)
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalln(err)
	}
	log.Println("Successfully applied migrations")
}
