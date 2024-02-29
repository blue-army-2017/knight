package view

import (
	"io"
	"time"

	"github.com/blue-army-2017/knight/model"
)

type SeasonGamesPage struct {
	Season *model.Season
	Games  []model.SeasonGame
	Flash  *Flash
}

func (p *SeasonGamesPage) Render(w io.Writer) {
	page := pages["season_games"]

	err := page.ExecuteTemplate(w, PAGE_TMPL, p)
	if err != nil {
		l.Error(err)
	}
}

type SeasonGamesNewPage struct {
	Season  *model.Season
	Game    *model.SeasonGame
	Members []model.Member
	Flash   *Flash
}

func (p *SeasonGamesNewPage) Render(w io.Writer) {
	page := pages["season_games_new"]

	if p.Game == nil {
		p.Game = &model.SeasonGame{
			Home: true,
			Date: time.Now().Format("2006-01-02"),
		}
	}

	err := page.ExecuteTemplate(w, PAGE_TMPL, p)
	if err != nil {
		l.Error(err)
	}
}
