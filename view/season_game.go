package view

import (
	"io"
	"slices"
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

func (p *SeasonGamesNewPage) IsMemberPresent(id string) bool {
	return slices.ContainsFunc(p.Game.PresentMembers, func(m model.Member) bool {
		return m.ID == id
	})
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

type SeasonGamesEditPage struct {
	Season  *model.Season
	Game    *model.SeasonGame
	Members []model.Member
	Flash   *Flash
}

func (p *SeasonGamesEditPage) IsMemberPresent(id string) bool {
	return slices.ContainsFunc(p.Game.PresentMembers, func(m model.Member) bool {
		return m.ID == id
	})
}

func (p *SeasonGamesEditPage) Render(w io.Writer) {
	page := pages["season_games_edit"]

	err := page.ExecuteTemplate(w, PAGE_TMPL, p)
	if err != nil {
		l.Error(err)
	}
}
