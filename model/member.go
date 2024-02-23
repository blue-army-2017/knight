package model

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	ID        string `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Active    bool
}

func FindAllMembers() (members []Member, err error) {
	result := db.
		Order("last_name").
		Order("first_name").
		Find(&members)
	err = result.Error
	return
}
