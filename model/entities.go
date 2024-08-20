package model

import (
	"time"

	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	ID        string `gorm:"primaryKey"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Active    bool   `gorm:"not null"`
}

type Season struct {
	gorm.Model
	ID    string `gorm:"primaryKey"`
	Name  string `gorm:"not null"`
	Games []SeasonGame
}

type SeasonGame struct {
	gorm.Model
	ID             string    `gorm:"primaryKey"`
	Opponent       string    `gorm:"not null"`
	Home           bool      `gorm:"not null"`
	Mode           string    `gorm:"not null"`
	Date           time.Time `gorm:"not null"`
	SeasonID       string    `gorm:"not null"`
	PresentMembers []Member  `gorm:"many2many:present_members;"`
}
