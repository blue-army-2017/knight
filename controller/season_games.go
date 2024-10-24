package controller

import (
	"context"
	"time"

	"github.com/blue-army-2017/knight/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type seasonGameRow struct {
	SeasonGame repository.SeasonGame
	SeasonName string
}

type SeasonGameDto struct {
	ID       string `form:"id"`
	Opponent string `form:"opponent"`
	Home     bool   `form:"home"`
	Mode     string `form:"mode"`
	Date     string `form:"date"`
	SeasonID string `form:"season_id"`
	Season   string
}

func CreateSeasonGameDto(game repository.SeasonGame, seasonName string) SeasonGameDto {
	return SeasonGameDto{
		ID:       game.ID,
		Opponent: game.Opponent,
		Home:     game.Home > 0.0,
		Mode:     game.Mode,
		Date:     game.Date,
		SeasonID: game.SeasonID,
		Season:   seasonName,
	}
}

func (dto *SeasonGameDto) ToModel() repository.SaveSeasonGameParams {
	var home float64
	if dto.Home {
		home = 1.0
	} else {
		home = 0.0
	}

	return repository.SaveSeasonGameParams{
		ID:       dto.ID,
		Opponent: dto.Opponent,
		Home:     home,
		Mode:     dto.Mode,
		Date:     dto.Date,
		SeasonID: dto.SeasonID,
	}
}

type SeasonGameController interface {
	GetIndex(seasonId string) Page
	GetNew() Page
	PostNew(game *SeasonGameDto) Page
	GetEdit(gameId string) Page
	PostEdit(game *SeasonGameDto, delete bool) Page
}

type DefaultSeasonGameController struct {
	repository repository.Querier
	ctx        context.Context
}

func NewSeasonGameController() SeasonGameController {
	return &DefaultSeasonGameController{
		repository: repository.New(db),
		ctx:        context.Background(),
	}
}

func (c *DefaultSeasonGameController) GetIndex(seasonId string) Page {
	var dtos []SeasonGameDto
	if len(seasonId) > 0 {
		rows, err := c.repository.FindAllSeasonGamesBySeason(c.ctx, seasonId)
		if err != nil {
			return &ErrorPage{
				Error: err,
			}
		}

		for _, row := range rows {
			dto := CreateSeasonGameDto(row.SeasonGame, row.SeasonName)
			dtos = append(dtos, dto)
		}
	} else {
		rows, err := c.repository.FindAllSeasonGames(c.ctx)
		if err != nil {
			return &ErrorPage{
				Error: err,
			}
		}

		for _, row := range rows {
			dto := CreateSeasonGameDto(row.SeasonGame, row.SeasonName)
			dtos = append(dtos, dto)
		}
	}

	return &HtmlPage{
		Template: "pages/games",
		Data: gin.H{
			"Games": dtos,
		},
	}
}

func (c *DefaultSeasonGameController) GetNew() Page {
	game := SeasonGameDto{
		ID:   uuid.NewString(),
		Home: true,
		Mode: "regular",
		Date: time.Now().Format("2006-01-02"),
	}

	seasonsData, err := c.repository.FindAllSeasons(c.ctx)
	if err != nil {
		return &ErrorPage{
			Error: err,
		}
	}
	var seasons []SeasonDto
	for _, data := range seasonsData {
		season := CreateSeasonDto(data)
		seasons = append(seasons, season)
	}

	return &HtmlPage{
		Template: "pages/games/new",
		Data: gin.H{
			"Game":    game,
			"Seasons": seasons,
		},
	}
}

func (c *DefaultSeasonGameController) PostNew(game *SeasonGameDto) Page {
	data := game.ToModel()
	err := c.repository.SaveSeasonGame(c.ctx, data)
	if err != nil {
		return &ErrorPage{
			Error: err,
		}
	}

	return &RedirectPage{
		Redirect: "/games",
	}
}

func (c *DefaultSeasonGameController) GetEdit(gameId string) Page {
	data, err := c.repository.FindSeasonGameById(c.ctx, gameId)
	if err != nil {
		return &ErrorPage{
			Error: err,
		}
	}
	game := CreateSeasonGameDto(data.SeasonGame, data.SeasonName)

	seasonsData, err := c.repository.FindAllSeasons(c.ctx)
	if err != nil {
		return &ErrorPage{
			Error: err,
		}
	}
	var seasons []SeasonDto
	for _, data := range seasonsData {
		season := CreateSeasonDto(data)
		seasons = append(seasons, season)
	}

	return &HtmlPage{
		Template: "pages/games/edit",
		Data: gin.H{
			"Game":    game,
			"Seasons": seasons,
		},
	}
}

func (c *DefaultSeasonGameController) PostEdit(game *SeasonGameDto, delete bool) Page {
	data := game.ToModel()

	var err error
	if delete {
		err = c.repository.DeleteSeasonGame(c.ctx, data.ID)
	} else {
		err = c.repository.SaveSeasonGame(c.ctx, data)
	}
	if err != nil {
		return &ErrorPage{
			Error: err,
		}
	}

	return &RedirectPage{
		Redirect: "/games",
	}
}
