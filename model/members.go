package model

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	ID        string `gorm:"primaryKey"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Active    bool   `gorm:"not null"`
}
