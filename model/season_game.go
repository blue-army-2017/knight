package model

import (
	"fmt"
	"regexp"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SeasonGame struct {
	gorm.Model
	ID             string `gorm:"primaryKey"`
	Opponent       string
	Home           bool
	Mode           string
	Date           string
	SeasonID       string
	PresentMembers []Member `gorm:"many2many:presence;"`
}

func FindAllSeasonGames(seasonId string) (games []SeasonGame, err error) {
	result := db.
		Model(&SeasonGame{}).
		Preload("PresentMembers").
		Order("date desc").
		Find(&games, "season_id = ?", seasonId)
	err = result.Error
	return
}

func FindSeasonGameByID(id string) (game SeasonGame, err error) {
	result := db.
		Model(&SeasonGame{}).
		Preload("PresentMembers").
		Find(&game, "id = ?", id)
	err = result.Error
	return
}

func (g *SeasonGame) Create() error {
	g.ID = uuid.NewString()
	if err := g.Validate(); err != nil {
		return err
	}

	result := db.Create(g)
	return result.Error
}

func (g *SeasonGame) Update() error {
	if err := g.Validate(); err != nil {
		return err
	}

	result := db.Save(g)
	if result.Error != nil {
		return result.Error
	}

	return db.
		Model(g).
		Association("PresentMembers").
		Replace(g.PresentMembers)
}

func (g *SeasonGame) Delete() error {
	err := db.Model(g).Association("PresentMembers").Clear()
	if err != nil {
		return err
	}

	result := db.Delete(g)
	return result.Error
}

func (g *SeasonGame) Validate() error {
	if len(g.ID) < 5 {
		return fmt.Errorf("ID must be at least 5 characters")
	}

	if len(g.Opponent) == 0 {
		return fmt.Errorf("Opponent must be set")
	}

	if g.Mode != "regular" && g.Mode != "playoffs" {
		return fmt.Errorf("Mode must be either 'regular' or 'playoffs'")
	}

	if matched, _ := regexp.MatchString(`^\d{4}-\d{2}-\d{2}$`, g.Date); !matched {
		return fmt.Errorf("Date must be in the format 'yyyy-MM-dd'")
	}

	if len(g.SeasonID) == 0 {
		return fmt.Errorf("Season ID must be set")
	}

	if len(g.PresentMembers) == 0 {
		return fmt.Errorf("At least one member must be present")
	}

	return nil
}
