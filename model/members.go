package model

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	ID        string `gorm:"primaryKey" form:"id"`
	FirstName string `gorm:"not null" form:"first_name"`
	LastName  string `gorm:"not null" form:"last_name"`
	Active    bool   `gorm:"not null" form:"active"`
}

type MemberRepository interface {
	FindAll() ([]Member, error)
	Create(member *Member) error
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

func (r *DefaultMemberRepository) Create(member *Member) error {
	result := db.Create(member)
	return result.Error
}
