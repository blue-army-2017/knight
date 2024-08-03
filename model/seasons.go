package model

import "gorm.io/gorm"

type Season struct {
	gorm.Model
	ID    string `gorm:"primaryKey"`
	Name  string `gorm:"not null"`
	Games []SeasonGame
}
