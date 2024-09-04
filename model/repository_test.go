package model

import (
	"testing"

	"github.com/onsi/gomega"
	"github.com/onsi/gomega/gstruct"
	"gorm.io/gorm"
)

func TestDefaultCRUDRepositoryFindAll(t *testing.T) {
	setupDB()
	defer teardownDB()
	tested := &DefaultCRUDRepository[Member]{}

	result, err := tested.FindAll("last_name,first_name")

	g := gomega.NewWithT(t)
	g.Expect(err).To(gomega.BeNil())
	g.Expect(result).To(gomega.HaveLen(5))
	g.Expect(result).To(gstruct.MatchAllElementsWithIndex(gstruct.IndexIdentity, gstruct.Elements{
		"0": gomega.HaveField("ID", "M004"),
		"1": gomega.HaveField("ID", "M005"),
		"2": gomega.HaveField("ID", "M001"),
		"3": gomega.HaveField("ID", "M003"),
		"4": gomega.HaveField("ID", "M002"),
	}))
}

func TestDefaultCRUDRepositoryFindAllBy(t *testing.T) {
	setupDB()
	defer teardownDB()
	tested := &DefaultCRUDRepository[Member]{}

	result, err := tested.FindAllBy("active", true, "last_name,first_name")

	g := gomega.NewWithT(t)
	g.Expect(err).To(gomega.BeNil())
	g.Expect(result).To(gomega.HaveLen(3))
	g.Expect(result).To(gstruct.MatchAllElementsWithIndex(gstruct.IndexIdentity, gstruct.Elements{
		"0": gomega.HaveField("ID", "M005"),
		"1": gomega.HaveField("ID", "M001"),
		"2": gomega.HaveField("ID", "M003"),
	}))
}

func TestDefaultCRUDRepositoryFindById(t *testing.T) {
	setupDB()
	defer teardownDB()
	tested := &DefaultCRUDRepository[Member]{}

	result, err := tested.FindById("M001")

	g := gomega.NewWithT(t)
	g.Expect(err).To(gomega.BeNil())
	g.Expect(*result).To(gomega.BeComparableTo(Member{
		ID:        "M001",
		FirstName: "John",
		LastName:  "Doe",
		Active:    true,
	}))
}

func TestDefaultCRUDRepositoryFindByIdNotFound(t *testing.T) {
	setupDB()
	defer teardownDB()
	tested := &DefaultCRUDRepository[Member]{}

	_, err := tested.FindById("xxx")

	g := gomega.NewWithT(t)
	g.Expect(err).To(gomega.MatchError(gorm.ErrRecordNotFound))
}

func TestDefaultCRUDRepositorySaveCreate(t *testing.T) {
	setupDB()
	defer teardownDB()
	tested := &DefaultCRUDRepository[Member]{}

	member := Member{
		ID:        "M006",
		FirstName: "Zaphod",
		LastName:  "Beeblebrox",
		Active:    true,
	}
	err := tested.Save(&member)

	g := gomega.NewWithT(t)
	g.Expect(err).To(gomega.BeNil())
	var entry Member
	if result := db.First(&entry, WHERE_ID_IS, member.ID); result.Error != nil {
		t.Fatal(result.Error)
	}
	g.Expect(entry).To(gomega.BeComparableTo(member))
}

func TestDefaultCRUDRepositorySaveUpdate(t *testing.T) {
	setupDB()
	defer teardownDB()
	tested := &DefaultCRUDRepository[Member]{}

	member := Member{
		ID:        "M003",
		FirstName: "Zaphod",
		LastName:  "Beeblebrox",
		Active:    true,
	}
	err := tested.Save(&member)

	g := gomega.NewWithT(t)
	g.Expect(err).To(gomega.BeNil())
	var entry Member
	if result := db.First(&entry, WHERE_ID_IS, member.ID); result.Error != nil {
		t.Fatal(result.Error)
	}
	g.Expect(entry).To(gomega.BeComparableTo(member))
}

func TestDefaultCRUDRepositoryDelete(t *testing.T) {
	setupDB()
	defer teardownDB()
	tested := &DefaultCRUDRepository[Member]{}

	member := Member{
		ID: "M002",
	}
	err := tested.Delete(&member)

	g := gomega.NewWithT(t)
	g.Expect(err).To(gomega.BeNil())
	var entry Member
	result := db.First(&entry, WHERE_ID_IS, member.ID)
	g.Expect(result.Error).To(gomega.MatchError(gorm.ErrRecordNotFound))
}
