package controller

import (
	"fmt"

	"github.com/blue-army-2017/knight/model"
	"github.com/gin-gonic/gin"
)

type GameStats struct {
	HomeGames  int
	AwayGames  int
	TotalGames int
}

type SeasonPresenceDto struct {
	Name           string
	MemberPresence []MemberPresenceDto
	GameStats
}

type MemberPresenceDto struct {
	Pos  int
	Name string
	GameStats
}

func CreatePresenceDto(seasonPresence []model.SeasonPresence, memberPresence []model.MemberPresence) []SeasonPresenceDto {
	memberPresenceBySeason := make(map[string][]model.MemberPresence)
	for _, member := range memberPresence {
		season := member.Season
		memberPresenceBySeason[season] = append(memberPresenceBySeason[season], member)
	}

	dto := []SeasonPresenceDto{}
	for _, season := range seasonPresence {
		seasonDto := SeasonPresenceDto{
			Name: season.Name,
			GameStats: GameStats{
				HomeGames:  season.HomeGames,
				AwayGames:  season.TotalGames - season.HomeGames,
				TotalGames: season.TotalGames,
			},
			MemberPresence: []MemberPresenceDto{},
		}

		for i, member := range memberPresenceBySeason[season.Name] {
			memberDto := MemberPresenceDto{
				Pos:  i + 1,
				Name: fmt.Sprintf("%s %s", member.FirstName, member.LastName),
				GameStats: GameStats{
					HomeGames:  member.HomeGames,
					AwayGames:  member.TotalGames - member.HomeGames,
					TotalGames: member.TotalGames,
				},
			}
			seasonDto.MemberPresence = append(seasonDto.MemberPresence, memberDto)
		}

		dto = append(dto, seasonDto)
	}
	return dto
}

type PresenceController interface {
	GetIndex() Page
}

type DefaultPresenceController struct {
	repository model.PresenceRepository
}

func NewPresenceController() PresenceController {
	return &DefaultPresenceController{
		repository: model.NewPresenceRepository(),
	}
}

func (c *DefaultPresenceController) GetIndex() Page {
	seasonPresence, err := c.repository.GetSeasonPresence()
	if err != nil {
		return &ErrorPage{
			Error: err,
		}
	}
	memberPresence, err := c.repository.GetMemberPresence()
	if err != nil {
		return &ErrorPage{
			Error: err,
		}
	}

	presence := CreatePresenceDto(seasonPresence, memberPresence)
	return &HtmlPage{
		Template: "pages/presence",
		Data: gin.H{
			"SeasonPresence": presence,
		},
	}
}
