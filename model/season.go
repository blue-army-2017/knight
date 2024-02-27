package model

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Season struct {
	gorm.Model
	ID    string `gorm:"primaryKey"`
	Name  string
	Games []SeasonGame
}

func FindAllSeasons() (seasons []Season, err error) {
	result := db.
		Order("created_at desc").
		Find(&seasons)
	err = result.Error
	return
}

func (s *Season) Create() error {
	s.ID = uuid.NewString()
	if err := s.Validate(); err != nil {
		return err
	}

	result := db.Create(s)
	return result.Error
}

func (s *Season) Validate() error {
	if err := uuid.Validate(s.ID); err != nil {
		return err
	}

	if len(s.Name) < 3 {
		return fmt.Errorf("Name must be at least 3 characters long")
	}

	return nil
}

type SeasonGame struct {
	gorm.Model
	ID             string `gorm:"primaryKey"`
	Opponent       string
	Home           bool
	Mode           string
	Date           string
	SeasonID       uint
	PresentMembers []Member `gorm:"many2many:presence;"`
}
