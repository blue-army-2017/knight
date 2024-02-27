package model

import (
	"gorm.io/gorm"
)

type Season struct {
	gorm.Model
	ID    string `gorm:"primaryKey"`
	Name  string
	Games []SeasonGame
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
