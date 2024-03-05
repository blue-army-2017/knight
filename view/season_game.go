package view

import (
	"html/template"
	"slices"

	"github.com/blue-army-2017/knight/model"
)

type SeasonGamesPage struct {
	Season *model.Season
	Games  []model.SeasonGame
	Flash  *Flash
}

func (p *SeasonGamesPage) Template() *template.Template {
	return pages["season_games"]
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

func (p *SeasonGamesNewPage) Template() *template.Template {
	return pages["season_games_new"]
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

func (p *SeasonGamesEditPage) Template() *template.Template {
	return pages["season_games_edit"]
}
