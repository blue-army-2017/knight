package controller

import (
	"context"
	"fmt"

	"github.com/blue-army-2017/knight/repository"
	"github.com/gin-gonic/gin"
)

type GameStats struct {
	HomeGames  int64
	AwayGames  int64
	TotalGames int64
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

func CreateSeasonPresenceDto(seasonPresence []repository.FindSeasonPresenceRow, memberPresence []repository.FindMemberPresenceRow) []SeasonPresenceDto {
	memberPresenceBySeason := make(map[string][]repository.FindMemberPresenceRow)
	for _, member := range memberPresence {
		season := member.Season
		memberPresenceBySeason[season] = append(memberPresenceBySeason[season], member)
	}

	dto := []SeasonPresenceDto{}
	for _, season := range seasonPresence {
		seasonDto := SeasonPresenceDto{
			Name: season.Name,
			GameStats: GameStats{
				HomeGames:  int64(season.HomeGames.Float64),
				AwayGames:  season.TotalGames - int64(season.HomeGames.Float64),
				TotalGames: season.TotalGames,
			},
			MemberPresence: []MemberPresenceDto{},
		}

		for _, member := range memberPresenceBySeason[season.Name] {
			memberDto := MemberPresenceDto{
				Name: fmt.Sprintf("%s %s", member.FirstName, member.LastName),
				GameStats: GameStats{
					HomeGames:  int64(member.HomeGames.Float64),
					AwayGames:  member.TotalGames - int64(member.HomeGames.Float64),
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
	repository repository.Querier
	ctx        context.Context
}

func NewPresenceController() PresenceController {
	return &DefaultPresenceController{
		repository: repository.New(db),
		ctx:        context.Background(),
	}
}

func (c *DefaultPresenceController) GetIndex() Page {
	seasonPresence, err := c.repository.FindSeasonPresence(c.ctx)
	if err != nil {
		return &ErrorPage{
			Error: err,
		}
	}
	memberPresence, err := c.repository.FindMemberPresence(c.ctx)
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
	game, err := c.repository.FindSeasonGameById(c.ctx, gameId)
	if err != nil {
		return &ErrorPage{
			Error: err,
		}
	}
	members, err := c.repository.FindAllMembers(c.ctx)
	if err != nil {
		return &ErrorPage{
			Error: err,
		}
	}

	presentMembers, err := c.repository.FindPresentMembersForGame(c.ctx, gameId)
	if err != nil {
		return &ErrorPage{
			Error: err,
		}
	}

	gameDto := CreateSeasonGameDto(game.SeasonGame, game.SeasonName)
	memberDtos := []MemberDto{}
	for _, member := range members {
		dto := CreateMemberDto(member)
		memberDtos = append(memberDtos, dto)
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
	err := c.repository.DeletePresentMembersForGame(c.ctx, gameId)
	if err != nil {
		return &ErrorPage{
			Error: err,
		}
	}

	for _, memberId := range presentMembers {
		err := c.repository.SavePresentMemberForGame(c.ctx, repository.SavePresentMemberForGameParams{
			SeasonGameID: gameId,
			MemberID:     memberId,
		})
		if err != nil {
			return &ErrorPage{
				Error: err,
			}
		}
	}

	return &RedirectPage{
		Redirect: "/presence",
	}
}
