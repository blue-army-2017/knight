package model

import (
	"time"

	"gorm.io/gorm"
)

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
