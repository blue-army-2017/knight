package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/blue-army-2017/knight/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	timeLocation *time.Location
)

func init() {
	var err error
	timeLocation, err = time.LoadLocation("Europe/Berlin")
	if err != nil {
		panic(err)
	}
}

func main() {
	args := os.Args
	dbFile := args[1]
	dataFile := args[2]

	migrate(dbFile, dataFile)
}

func migrate(dbFile, dataFile string) {
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	checkErr(err)

	err = db.AutoMigrate(&model.Member{})
	checkErr(err)
	err = db.AutoMigrate(&model.Season{})
	checkErr(err)
	err = db.AutoMigrate(&model.SeasonGame{})
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
	// Members
	result := db.
		Clauses(clause.OnConflict{UpdateAll: true}).
		Create(&members)
	checkErr(result.Error)
	log.Printf("Migrated %d members\n", result.RowsAffected)

	var seasons []model.Season
	var seasonGames []model.SeasonGame
	for seasonId, seasonData := range data.Season {
		season := model.Season{
			ID:      seasonId,
			Name:    seasonData.Name,
			Created: seasonData.Created.In(timeLocation).Format("2006-01-02"),
		}
		seasons = append(seasons, season)

		for gameId, gameData := range seasonData.Games {
			game := model.SeasonGame{
				ID:       gameId,
				Opponent: gameData.Opponent,
				Home:     gameData.Home,
				Mode:     gameData.Mode,
				Date:     gameData.Date.In(timeLocation).Format("2006-01-02"),
				SeasonID: seasonId,
			}

			var presentMembers []model.Member
			for _, memberId := range gameData.PresentMembers {
				presentMembers = append(presentMembers, model.Member{ID: memberId})
			}
			game.PresentMembers = presentMembers

			seasonGames = append(seasonGames, game)
		}
	}
	// Seasons
	result = db.
		Clauses(clause.OnConflict{UpdateAll: true}).
		Create(&seasons)
	checkErr(result.Error)
	log.Printf("Migrated %d seasons\n", result.RowsAffected)
	// Season Games
	result = db.
		Clauses(clause.OnConflict{UpdateAll: true}).
		Create(&seasonGames)
	checkErr(result.Error)
	log.Printf("Migrated %d season games\n", result.RowsAffected)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}
