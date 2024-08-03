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
	FindById(id string) (*Member, error)
	Create(member *Member) error
	Update(member *Member) error
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

func (r *DefaultMemberRepository) FindById(id string) (*Member, error) {
	member := Member{
		ID: id,
	}
	result := db.First(&member)
	return &member, result.Error
}

func (r *DefaultMemberRepository) Create(member *Member) error {
	result := db.Create(member)
	return result.Error
}

func (r *DefaultMemberRepository) Update(member *Member) error {
	result := db.Save(member)
	return result.Error
}
