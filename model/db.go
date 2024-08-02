package model

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("knight.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}

	models := []any{
		Member{},
	}
	for _, model := range models {
		if err := db.AutoMigrate(&model); err != nil {
			log.Fatalln(err.Error())
		}
	}
}
