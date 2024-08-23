package controller

import (
	"github.com/blue-army-2017/knight/model"
	"github.com/gin-gonic/gin"
)

type SeasonGameDto struct {
	ID       string `form:"id"`
	Opponent string `form:"opponent"`
	Home     bool   `form:"home"`
	Mode     string `form:"mode"`
	Date     string `form:"date"`
	SeasonID string `form:"season_id"`
	Season   string
}

func CreateSeasonGameDto(game *model.SeasonGame) *SeasonGameDto {
	return &SeasonGameDto{
		ID:       game.ID,
		Opponent: game.Opponent,
		Home:     game.Home,
		Mode:     game.Mode,
		Date:     game.Date,
		SeasonID: game.SeasonID,
		Season:   game.Season.Name,
	}
}

func (dto *SeasonGameDto) ToModel() *model.SeasonGame {
	return &model.SeasonGame{
		ID:       dto.ID,
		Opponent: dto.Opponent,
		Home:     dto.Home,
		Mode:     dto.Mode,
		Date:     dto.Date,
		SeasonID: dto.SeasonID,
	}
}

type SeasonGameController interface {
	GetIndex() Page
}

type DefaultSeasonGameController struct {
	repository model.CRUDRepository[model.SeasonGame]
}

func NewSeasonGameController() SeasonGameController {
	return &DefaultSeasonGameController{
		repository: model.NewCRUDRepository[model.SeasonGame](),
	}
}

func (c *DefaultSeasonGameController) GetIndex() Page {
	games, err := c.repository.FindAll("date desc")
	if err != nil {
		return &ErrorPage{
			Error: err,
		}
	}

	var dtos []SeasonGameDto
	for _, game := range games {
		dto := CreateSeasonGameDto(&game)
		dtos = append(dtos, *dto)
	}

	return &HtmlPage{
		Template: "pages/games",
		Data: gin.H{
			"Games": dtos,
		},
	}
}
