package controller

import (
	"testing"

	"github.com/blue-army-2017/knight/model"
	"github.com/google/uuid"
	"github.com/onsi/gomega"
)

func TestCreateMemberDto(t *testing.T) {
	id := uuid.NewString()
	member := model.Member{
		ID:        id,
		FirstName: "Arthur",
		LastName:  "Dent",
		Active:    true,
	}
	expected := MemberDto{
		ID:        id,
		FirstName: "Arthur",
		LastName:  "Dent",
		Active:    true,
	}

	result := CreateMemberDto(&member)

	g := gomega.NewWithT(t)
	g.Expect(*result).To(gomega.Equal(expected))
}

func TestMemberDtoToModel(t *testing.T) {
	id := uuid.NewString()
	member := MemberDto{
		ID:        id,
		FirstName: "Arthur",
		LastName:  "Dent",
		Active:    true,
	}
	expected := model.Member{
		ID:        id,
		FirstName: "Arthur",
		LastName:  "Dent",
		Active:    true,
	}

	result := member.ToModel()

	g := gomega.NewWithT(t)
	g.Expect(*result).To(gomega.Equal(expected))
}
