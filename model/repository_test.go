package model

import (
	"testing"

	"github.com/onsi/gomega"
	"github.com/onsi/gomega/gstruct"
	"gorm.io/gorm"
)

const (
	WHERE_ID = "id = ?"
)

var members = []Member{
	{
		ID:        "00",
		FirstName: "Marvin",
		LastName:  "The Paranoid Android",
		Active:    false,
	},
	{
		ID:        "42",
		FirstName: "Arthur",
		LastName:  "Dent",
		Active:    true,
	},
	{
		ID:        "43",
		FirstName: "Ford",
		LastName:  "Prefect",
		Active:    true,
	},
}

func setupRepoTest() *DefaultCRUDRepository[Member] {
	r := db.Create(&members)
	if r.Error != nil {
		panic(r.Error)
	}

	return &DefaultCRUDRepository[Member]{}
}

func teardownRepoTest() {
	r := db.Exec("DELETE FROM members")
	if r.Error != nil {
		panic(r.Error)
	}
}

func TestDefaultCRUDRepositoryFindAll(t *testing.T) {
	tested := setupRepoTest()
	defer teardownRepoTest()

	result, err := tested.FindAll("last_name", "first_name")

	g := gomega.NewWithT(t)
	g.Expect(err).To(gomega.BeNil())
	g.Expect(result).To(gomega.HaveLen(3))
	g.Expect(result).To(gstruct.MatchAllElementsWithIndex(gstruct.IndexIdentity, gstruct.Elements{
		"0": gomega.BeComparableTo(members[1]),
		"1": gomega.BeComparableTo(members[2]),
		"2": gomega.BeComparableTo(members[0]),
	}))
}

func TestDefaultCRUDRepositoryFindById(t *testing.T) {
	tested := setupRepoTest()
	defer teardownRepoTest()

	result, err := tested.FindById("42")

	g := gomega.NewWithT(t)
	g.Expect(err).To(gomega.BeNil())
	g.Expect(*result).To(gomega.BeComparableTo(members[1]))
}

func TestDefaultCRUDRepositoryFindByIdNotFound(t *testing.T) {
	tested := setupRepoTest()
	defer teardownRepoTest()

	_, err := tested.FindById("xxx")

	g := gomega.NewWithT(t)
	g.Expect(err).To(gomega.MatchError(gorm.ErrRecordNotFound))
}

func TestDefaultCRUDRepositorySaveCreate(t *testing.T) {
	tested := setupRepoTest()
	defer teardownRepoTest()

	member := Member{
		FirstName: "Zaphod",
		LastName:  "Beeblebrox",
		Active:    true,
	}
	err := tested.Save(&member)

	g := gomega.NewWithT(t)
	g.Expect(err).To(gomega.BeNil())
	var entry Member
	if result := db.First(&entry, WHERE_ID, member.ID); result.Error != nil {
		t.Fatal(result.Error)
	}
	g.Expect(entry).To(gomega.BeComparableTo(member))
}

func TestDefaultCRUDRepositorySaveUpdate(t *testing.T) {
	tested := setupRepoTest()
	defer teardownRepoTest()

	member := Member{
		ID:        "43",
		FirstName: "Zaphod",
		LastName:  "Beeblebrox",
		Active:    true,
	}
	err := tested.Save(&member)

	g := gomega.NewWithT(t)
	g.Expect(err).To(gomega.BeNil())
	var entry Member
	if result := db.First(&entry, WHERE_ID, member.ID); result.Error != nil {
		t.Fatal(result.Error)
	}
	g.Expect(entry).To(gomega.BeComparableTo(member))
}

func TestDefaultCRUDRepositoryDelete(t *testing.T) {
	tested := setupRepoTest()
	defer teardownRepoTest()

	member := Member{
		ID: "42",
	}
	err := tested.Delete(&member)

	g := gomega.NewWithT(t)
	g.Expect(err).To(gomega.BeNil())
	var entry Member
	result := db.First(&entry, WHERE_ID, member.ID)
	g.Expect(result.Error).To(gomega.MatchError(gorm.ErrRecordNotFound))
}
