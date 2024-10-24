package controller

import (
	"time"

	"github.com/blue-army-2017/knight/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

type SeasonDto struct {
	ID      string `form:"id"`
	Name    string `form:"name"`
	Created string `form:"created"`
}

func CreateSeasonDto(season repository.Season) SeasonDto {
	return SeasonDto{
		ID:      season.ID,
		Name:    season.Name,
		Created: season.Created,
	}
}

func (dto *SeasonDto) ToModel() repository.SaveSeasonParams {
	return repository.SaveSeasonParams{
		ID:      dto.ID,
		Name:    dto.Name,
		Created: dto.Created,
	}
}

type SeasonController interface {
	GetIndex() Page
	GetNew() Page
	PostNew(season *SeasonDto) Page
	GetEdit(id string) Page
	PostEdit(season *SeasonDto, delete bool) Page
}

type DefaultSeasonController struct {
	repository repository.Querier
	ctx        context.Context
}

func NewSeasonController() SeasonController {
	return &DefaultSeasonController{
		repository: repository.New(db),
		ctx:        context.Background(),
	}
}

func (c *DefaultSeasonController) GetIndex() Page {
	seasons, err := c.repository.FindAllSeasons(c.ctx)
	if err != nil {
		return &ErrorPage{
			Error: err,
		}
	}

	var dtos []SeasonDto
	for _, season := range seasons {
		dto := CreateSeasonDto(season)
		dtos = append(dtos, dto)
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
		ID:      uuid.NewString(),
		Created: time.Now().Format("2006-01-02"),
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
	err := c.repository.SaveSeason(c.ctx, data)
	if err != nil {
		return &ErrorPage{
			Error: err,
		}
	}

	return &RedirectPage{
		Redirect: "/seasons",
	}
}

func (c *DefaultSeasonController) GetEdit(id string) Page {
	season, err := c.repository.FindSeasonById(c.ctx, id)
	if err != nil {
		return &ErrorPage{
			Error: err,
		}
	}

	return &HtmlPage{
		Template: "pages/seasons/edit",
		Data: gin.H{
			"Season": CreateSeasonDto(season),
		},
	}

}

func (c *DefaultSeasonController) PostEdit(season *SeasonDto, delete bool) Page {
	data := season.ToModel()

	var err error
	if delete {
		err = c.repository.DeleteSeason(c.ctx, data.ID)
	} else {
		err = c.repository.SaveSeason(c.ctx, data)
	}
	if err != nil {
		return &ErrorPage{
			Error: err,
		}
	}

	return &RedirectPage{
		Redirect: "/seasons",
	}
}
