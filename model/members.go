package model

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	ID        string `gorm:"primaryKey"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Active    bool   `gorm:"not null"`
}

type MemberRepository interface {
	FindAll() ([]Member, error)
}

type DefaultMemberRepository struct{}

func NewMemberRepository() MemberRepository {
	return &DefaultMemberRepository{}
}

func (r *DefaultMemberRepository) FindAll() ([]Member, error) {
	var members []Member
	result := db.
		Order("last_name").
		Order("first_name").
		Find(&members)
	return members, result.Error
}
