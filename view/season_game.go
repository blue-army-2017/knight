package view

import (
	"io"

	"github.com/blue-army-2017/knight/model"
)

type SeasonGamesPage struct {
	SeasonID string
	Games    []model.SeasonGame
	Flash    *Flash
}

func (p *SeasonGamesPage) Render(w io.Writer) {
	page := pages["season_games"]

	err := page.ExecuteTemplate(w, PAGE_TMPL, p)
	if err != nil {
		l.Error(err)
	}
}
