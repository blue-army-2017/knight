package view

import (
	"html/template"

	"github.com/blue-army-2017/knight/model"
)

type SeasonsPage struct {
	Seasons []model.Season
	Flash   *Flash
}

func (p *SeasonsPage) Template() *template.Template {
	return pages["seasons"]
}

type SeasonsShowPage struct {
	Season *model.Season
	Flash  *Flash
}

func (p *SeasonsShowPage) Template() *template.Template {
	return pages["seasons_show"]
}

type SeasonsNewPage struct {
	Season *model.Season
	Flash  *Flash
}

func (p *SeasonsNewPage) Template() *template.Template {
	return pages["seasons_new"]
}

type SeasonsEditPage struct {
	Season *model.Season
	Flash  *Flash
}

func (p *SeasonsEditPage) Template() *template.Template {
	return pages["seasons_edit"]
}
