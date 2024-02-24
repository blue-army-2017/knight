package model

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

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

func FindMemberByID(id string) (member Member, err error) {
	result := db.First(&member, "id = ?", id)
	err = result.Error
	return
}

func (m *Member) Create() error {
	m.ID = uuid.NewString()
	if err := m.Validate(); err != nil {
		return err
	}

	result := db.Create(m)
	return result.Error
}

func (m *Member) Update() error {
	if err := m.Validate(); err != nil {
		return err
	}

	result := db.Save(m)
	return result.Error
}

func (m *Member) Delete() error {
	result := db.Delete(m)
	return result.Error
}

func (m *Member) Validate() error {
	if err := uuid.Validate(m.ID); err != nil {
		return err
	}

	if len(m.FirstName) == 0 {
		return fmt.Errorf("First name must be set")
	}

	if len(m.LastName) == 0 {
		return fmt.Errorf("Last name must be set")
	}

	return nil
}
