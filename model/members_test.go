package model

import (
	"testing"

	"github.com/google/uuid"
	"github.com/onsi/gomega"
	"github.com/onsi/gomega/gstruct"
	"gorm.io/gorm"
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

func setupMembersTest() *DefaultMemberRepository {
	r := db.Create(&members)
	if r.Error != nil {
		panic(r.Error)
	}

	return &DefaultMemberRepository{}
}

func teardownMembersTest() {
	r := db.Exec("DELETE FROM members")
	if r.Error != nil {
		panic(r.Error)
	}
}

func TestDefaultMemberRepositoryFindAll(t *testing.T) {
	tested := setupMembersTest()
	defer teardownMembersTest()

	result, err := tested.FindAll()

	g := gomega.NewWithT(t)
	g.Expect(err).To(gomega.BeNil())
	g.Expect(result).To(gomega.HaveLen(3))
	g.Expect(result).To(gstruct.MatchAllElementsWithIndex(gstruct.IndexIdentity, gstruct.Elements{
		"0": gomega.BeComparableTo(members[1]),
		"1": gomega.BeComparableTo(members[2]),
		"2": gomega.BeComparableTo(members[0]),
	}))
}

func TestDefaultMemberRepositoryFindById(t *testing.T) {
	tested := setupMembersTest()
	defer teardownMembersTest()

	result, err := tested.FindById("42")

	g := gomega.NewWithT(t)
	g.Expect(err).To(gomega.BeNil())
	g.Expect(*result).To(gomega.BeComparableTo(members[1]))
}

func TestDefaultMemberRepositoryFindByIdNotFound(t *testing.T) {
	tested := setupMembersTest()
	defer teardownMembersTest()

	_, err := tested.FindById("xxx")

	g := gomega.NewWithT(t)
	g.Expect(err).To(gomega.MatchError(gorm.ErrRecordNotFound))
}

func TestDefaultMemberRepositoryCreate(t *testing.T) {
	tested := setupMembersTest()
	defer teardownMembersTest()

	member := Member{
		FirstName: "Zaphod",
		LastName:  "Beeblebrox",
		Active:    true,
	}
	err := tested.Create(&member)

	g := gomega.NewWithT(t)
	g.Expect(err).To(gomega.BeNil())
	g.Expect(member.ID).To(gomega.HaveLen(len(uuid.NewString())))
	var entry Member
	if result := db.First(&entry, "id = ?", member.ID); result.Error != nil {
		t.Fatal(result.Error)
	}
	g.Expect(entry).To(gomega.BeComparableTo(member))
}

func TestDefaultMemberRepositoryUpdate(t *testing.T) {
	tested := setupMembersTest()
	defer teardownMembersTest()

	member := Member{
		ID:        "43",
		FirstName: "Zaphod",
		LastName:  "Beeblebrox",
		Active:    true,
	}
	err := tested.Update(&member)

	g := gomega.NewWithT(t)
	g.Expect(err).To(gomega.BeNil())
	var entry Member
	if result := db.First(&entry, "id = ?", member.ID); result.Error != nil {
		t.Fatal(result.Error)
	}
	g.Expect(entry).To(gomega.BeComparableTo(member))
}
