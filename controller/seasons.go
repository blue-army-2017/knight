package controller

import (
	"github.com/blue-army-2017/knight/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SeasonDto struct {
	ID   string `form:"id"`
	Name string `form:"name"`
}

func CreateSeasonDto(season *model.Season) *SeasonDto {
	return &SeasonDto{
		ID:   season.ID,
		Name: season.Name,
	}
}

func (dto *SeasonDto) ToModel() *model.Season {
	return &model.Season{
		ID:   dto.ID,
		Name: dto.Name,
	}
}

type SeasonController interface {
	GetIndex() Page
	GetNew() Page
	PostNew(season *SeasonDto) Page
}

type DefaultSeasonController struct {
	repository model.CRUDRepository[model.Season]
}

func NewSeasonController() SeasonController {
	return &DefaultSeasonController{
		repository: model.NewCRUDRepository[model.Season](),
	}
}

func (c *DefaultSeasonController) GetIndex() Page {
	seasons, err := c.repository.FindAll("created_at desc")
	if err != nil {
		return &ErrorPage{
			Error: err,
		}
	}

	var dtos []SeasonDto
	for _, season := range seasons {
		dto := CreateSeasonDto(&season)
		dtos = append(dtos, *dto)
	}

	return &HtmlPage{
		Template: "pages/seasons",
		Data: gin.H{
			"Seasons": dtos,
		},
	}
}

func (c *DefaultSeasonController) GetNew() Page {
	season := SeasonDto{
		ID: uuid.NewString(),
	}

	return &HtmlPage{
		Template: "pages/seasons/new",
		Data: gin.H{
			"Season": season,
		},
	}
}

func (c *DefaultSeasonController) PostNew(season *SeasonDto) Page {
	data := season.ToModel()
	err := c.repository.Create(data)
	if err != nil {
		return &ErrorPage{
			Error: err,
		}
	}

	return &RedirectPage{
		Redirect: "/seasons",
	}
}
