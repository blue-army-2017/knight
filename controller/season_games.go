package controller

import (
	"time"

	"github.com/blue-army-2017/knight/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	GetIndex(seasonId string) Page
	GetNew() Page
	PostNew(game *SeasonGameDto) Page
	GetEdit(gameId string) Page
	PostEdit(game *SeasonGameDto, delete bool) Page
}

type DefaultSeasonGameController struct {
	gamesRepository   model.CRUDRepository[model.SeasonGame]
	seasonsRepository model.CRUDRepository[model.Season]
}

func NewSeasonGameController() SeasonGameController {
	return &DefaultSeasonGameController{
		gamesRepository:   model.NewCRUDRepository[model.SeasonGame](),
		seasonsRepository: model.NewCRUDRepository[model.Season](),
	}
}

func (c *DefaultSeasonGameController) GetIndex(seasonId string) Page {
	var games []model.SeasonGame
	var err error
	if len(seasonId) > 0 {
		games, err = c.gamesRepository.FindAllBy("season_id", seasonId, "date desc")
	} else {
		games, err = c.gamesRepository.FindAll("date desc")
	}
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

func (c *DefaultSeasonGameController) GetNew() Page {
	game := SeasonGameDto{
		ID:   uuid.NewString(),
		Home: true,
		Mode: "regular",
		Date: time.Now().Format("2006-01-02"),
	}

	seasonsData, err := c.seasonsRepository.FindAll("created desc")
	if err != nil {
		return &ErrorPage{
			Error: err,
		}
	}
	var seasons []SeasonDto
	for _, data := range seasonsData {
		season := CreateSeasonDto(&data)
		seasons = append(seasons, *season)
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
	err := c.gamesRepository.Save(data)
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
	data, err := c.gamesRepository.FindById(gameId)
	if err != nil {
		return &ErrorPage{
			Error: err,
		}
	}
	game := CreateSeasonGameDto(data)

	seasonsData, err := c.seasonsRepository.FindAll("created desc")
	if err != nil {
		return &ErrorPage{
			Error: err,
		}
	}
	var seasons []SeasonDto
	for _, data := range seasonsData {
		season := CreateSeasonDto(&data)
		seasons = append(seasons, *season)
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
		err = c.gamesRepository.Delete(data)
	} else {
		err = c.gamesRepository.Save(data)
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
