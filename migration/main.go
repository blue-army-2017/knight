package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/blue-army-2017/knight/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func main() {
	args := os.Args
	dbFile := args[1]
	dataFile := args[2]

	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	checkErr(err)

	err = db.AutoMigrate(&model.Member{})
	checkErr(err)

	dataFileContent, err := os.ReadFile(dataFile)
	checkErr(err)
	var data LegacyData
	err = json.Unmarshal(dataFileContent, &data)
	checkErr(err)

	var members []model.Member
	for id, memberData := range data.Member {
		member := model.Member{
			ID:        id,
			FirstName: memberData.FirstName,
			LastName:  memberData.LastName,
			Active:    memberData.Active,
		}
		members = append(members, member)
	}
	result := db.
		Clauses(clause.OnConflict{UpdateAll: true}).
		Create(&members)
	checkErr(result.Error)
	log.Printf("Migrated %d members\n", result.RowsAffected)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}
