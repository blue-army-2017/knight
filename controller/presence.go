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
	Name string
	GameStats
}

func CreateSeasonPresenceDto(seasonPresence []model.SeasonPresence, memberPresence []model.MemberPresence) []SeasonPresenceDto {
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

		for _, member := range memberPresenceBySeason[season.Name] {
			memberDto := MemberPresenceDto{
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
	GetEdit(gameId string) Page
	GetEditPost(gameId string, presentMembers []string) Page
}

type DefaultPresenceController struct {
	presenceRepository   model.PresenceRepository
	seasonGameRepository model.CRUDRepository[model.SeasonGame]
	memberRepository     model.CRUDRepository[model.Member]
}

func NewPresenceController() PresenceController {
	return &DefaultPresenceController{
		presenceRepository:   model.NewPresenceRepository(),
		seasonGameRepository: model.NewCRUDRepository[model.SeasonGame](),
		memberRepository:     model.NewCRUDRepository[model.Member](),
	}
}

func (c *DefaultPresenceController) GetIndex() Page {
	seasonPresence, err := c.presenceRepository.GetSeasonPresence()
	if err != nil {
		return &ErrorPage{
			Error: err,
		}
	}
	memberPresence, err := c.presenceRepository.GetMemberPresence()
	if err != nil {
		return &ErrorPage{
			Error: err,
		}
	}

	presence := CreateSeasonPresenceDto(seasonPresence, memberPresence)
	return &HtmlPage{
		Template: "pages/presence",
		Data: gin.H{
			"SeasonPresence": presence,
		},
	}
}

func (c *DefaultPresenceController) GetEdit(gameId string) Page {
	game, err := c.seasonGameRepository.FindById(gameId)
	if err != nil {
		return &ErrorPage{
			Error: err,
		}
	}
	members, err := c.memberRepository.FindAll("last_name, first_name")
	if err != nil {
		return &ErrorPage{
			Error: err,
		}
	}

	presentMembers := []string{}
	for _, member := range game.PresentMembers {
		presentMembers = append(presentMembers, member.ID)
	}

	gameDto := CreateSeasonGameDto(game)
	memberDtos := []MemberDto{}
	for _, member := range members {
		dto := CreateMemberDto(&member)
		memberDtos = append(memberDtos, *dto)
	}

	return &HtmlPage{
		Template: "pages/presence/edit",
		Data: gin.H{
			"Game":           gameDto,
			"Members":        memberDtos,
			"PresentMembers": presentMembers,
		},
	}
}

func (c *DefaultPresenceController) GetEditPost(gameId string, presentMembers []string) Page {
	if err := c.presenceRepository.SavePresentMembers(gameId, presentMembers); err != nil {
		return &ErrorPage{
			Error: err,
		}
	}

	return &RedirectPage{
		Redirect: "/presence",
	}
}
