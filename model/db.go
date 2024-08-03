package model

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dbName, dbNameExists := os.LookupEnv("DB_NAME")
	if !dbNameExists {
		dbName = "knight.db"
	}

	var err error
	db, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}

	models := []any{
		Member{},
		Season{},
		SeasonGame{},
	}
	for _, model := range models {
		if err := db.AutoMigrate(&model); err != nil {
			log.Fatalln(err.Error())
		}
	}
}
