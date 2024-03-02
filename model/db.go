package model

import (
	"github.com/blue-army-2017/knight/util"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open(util.Config.DB), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Member{})
	db.AutoMigrate(&Season{})
	db.AutoMigrate(&SeasonGame{})
}
