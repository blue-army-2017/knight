package model

import (
	"testing"

	"github.com/onsi/gomega"
	"github.com/onsi/gomega/gstruct"
)

func TestDefaultPresenceRepositoryGetSeasonPresence(t *testing.T) {
	setupDB()
	defer teardownDB()
	tested := &DefaultPresenceRepository{}

	result, err := tested.GetSeasonPresence()

	g := gomega.NewWithT(t)
	g.Expect(err).To(gomega.BeNil())
	g.Expect(result).To(gomega.HaveLen(3))
	g.Expect(result).To(gstruct.MatchAllElementsWithIndex(gstruct.IndexIdentity, gstruct.Elements{
		"0": gomega.BeComparableTo(SeasonPresence{
			Name:       SEASON_FALL,
			HomeGames:  3,
			TotalGames: 5,
		}),
		"1": gomega.BeComparableTo(SeasonPresence{
			Name:       SEASON_SUMMER,
			HomeGames:  3,
			TotalGames: 5,
		}),
		"2": gomega.BeComparableTo(SeasonPresence{
			Name:       SEASON_SPRING,
			HomeGames:  3,
			TotalGames: 5,
		}),
	}))
}

func TestDefaultPresenceRepositoryGetMemberPresence(t *testing.T) {
	setupDB()
	defer teardownDB()
	tested := &DefaultPresenceRepository{}

	result, err := tested.GetMemberPresence()

	g := gomega.NewWithT(t)
	g.Expect(err).To(gomega.BeNil())
	g.Expect(result).To(gomega.HaveLen(9))
	g.Expect(result).To(gstruct.MatchAllElementsWithIndex(gstruct.IndexIdentity, gstruct.Elements{
		"0": gomega.BeComparableTo(MemberPresence{
			Season:     SEASON_FALL,
			LastName:   "Davis",
			FirstName:  "Charlie",
			HomeGames:  0,
			TotalGames: 2,
		}),
		"1": gomega.BeComparableTo(MemberPresence{
			Season:     SEASON_FALL,
			LastName:   "Doe",
			FirstName:  "John",
			HomeGames:  0,
			TotalGames: 2,
		}),
		"2": gomega.BeComparableTo(MemberPresence{
			Season:     SEASON_FALL,
			LastName:   "Johnson",
			FirstName:  "Alice",
			HomeGames:  2,
			TotalGames: 2,
		}),
		"3": gomega.BeComparableTo(MemberPresence{
			Season:     SEASON_SUMMER,
			LastName:   "Johnson",
			FirstName:  "Alice",
			HomeGames:  3,
			TotalGames: 3,
		}),
		"4": gomega.BeComparableTo(MemberPresence{
			Season:     SEASON_SUMMER,
			LastName:   "Davis",
			FirstName:  "Charlie",
			HomeGames:  1,
			TotalGames: 2,
		}),
		"5": gomega.BeComparableTo(MemberPresence{
			Season:     SEASON_SUMMER,
			LastName:   "Doe",
			FirstName:  "John",
			HomeGames:  1,
			TotalGames: 2,
		}),
		"6": gomega.BeComparableTo(MemberPresence{
			Season:     SEASON_SPRING,
			LastName:   "Davis",
			FirstName:  "Charlie",
			HomeGames:  2,
			TotalGames: 2,
		}),
		"7": gomega.BeComparableTo(MemberPresence{
			Season:     SEASON_SPRING,
			LastName:   "Doe",
			FirstName:  "John",
			HomeGames:  1,
			TotalGames: 2,
		}),
		"8": gomega.BeComparableTo(MemberPresence{
			Season:     SEASON_SPRING,
			LastName:   "Johnson",
			FirstName:  "Alice",
			HomeGames:  1,
			TotalGames: 2,
		}),
	}))
}

func TestDefaultPresenceRepositorySavePresentMembers(t *testing.T) {
	setupDB()
	defer teardownDB()
	tested := &DefaultPresenceRepository{}

	gameId := "G001"
	presentMembers := []string{"M002", "M003", "M005"}
	err := tested.SavePresentMembers(gameId, presentMembers)

	g := gomega.NewWithT(t)
	g.Expect(err).To(gomega.BeNil())
	var memberIds []string
	if result := db.Raw("SELECT member_id FROM present_members WHERE season_game_id = ?", gameId).Scan(&memberIds); result.Error != nil {
		t.Fatal(result.Error)
	}
	g.Expect(memberIds).To(gomega.HaveLen(3))
	g.Expect(memberIds).To(gomega.ContainElements(presentMembers))
}
