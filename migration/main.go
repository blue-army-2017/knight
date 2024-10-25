package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/blue-army-2017/knight/repository"
	_ "github.com/mattn/go-sqlite3"
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
	db, err := sql.Open("sqlite3", dbFile)
	checkErr(err)
	repo := repository.New(db)
	ctx := context.Background()

	dataFileContent, err := os.ReadFile(dataFile)
	checkErr(err)
	var data LegacyData
	err = json.Unmarshal(dataFileContent, &data)
	checkErr(err)

	// Members
	for id, memberData := range data.Member {
		var active float64
		if memberData.Active {
			active = 1.0
		} else {
			active = 0.0
		}

		member := repository.SaveMemberParams{
			ID:        id,
			FirstName: memberData.FirstName,
			LastName:  memberData.LastName,
			Active:    active,
		}
		err := repo.SaveMember(ctx, member)
		checkErr(err)
	}
	log.Printf("Migrated %d members\n", len(data.Member))

	// Seasons
	for seasonId, seasonData := range data.Season {
		season := repository.SaveSeasonParams{
			ID:      seasonId,
			Name:    seasonData.Name,
			Created: seasonData.Created.In(timeLocation).Format("2006-01-02"),
		}
		err := repo.SaveSeason(ctx, season)
		checkErr(err)
		log.Printf("Migrated season %s\n", seasonData.Name)

		// Season Games
		for gameId, gameData := range seasonData.Games {
			var home float64
			if gameData.Home {
				home = 1.0
			} else {
				home = 0.0
			}

			game := repository.SaveSeasonGameParams{
				ID:       gameId,
				Opponent: gameData.Opponent,
				Home:     home,
				Mode:     gameData.Mode,
				Date:     gameData.Date.In(timeLocation).Format("2006-01-02"),
				SeasonID: seasonId,
			}
			err := repo.SaveSeasonGame(ctx, game)
			checkErr(err)

			// Present Members
			err = repo.DeletePresentMembersForGame(ctx, gameId)
			checkErr(err)

			for _, memberId := range gameData.PresentMembers {
				err := repo.SavePresentMemberForGame(ctx, repository.SavePresentMemberForGameParams{
					SeasonGameID: gameId,
					MemberID:     memberId,
				})
				checkErr(err)
			}
		}
		log.Printf("Migrated %d games for season %s\n", len(seasonData.Games), seasonData.Name)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}
