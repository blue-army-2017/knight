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

func FindSeasonByID(id string) (season Season, err error) {
	result := db.Find(&season, "id = ?", id)
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

func (s *Season) Update() error {
	if err := s.Validate(); err != nil {
		return err
	}

	result := db.Save(s)
	return result.Error
}

func (s *Season) Delete() error {
	result := db.Delete(s)
	return result.Error
}

func (s *Season) Validate() error {
	if len(s.ID) < 5 {
		return fmt.Errorf("ID must be at least 5 characters")
	}

	if len(s.Name) < 3 {
		return fmt.Errorf("Name must be at least 3 characters")
	}

	return nil
}
