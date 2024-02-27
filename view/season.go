package view

import (
	"io"

	"github.com/blue-army-2017/knight/model"
)

type SeasonsPage struct {
	Seasons []model.Season
	Flash   *Flash
}

func (p *SeasonsPage) Render(w io.Writer) {
	page := pages["seasons"]

	err := page.ExecuteTemplate(w, PAGE_TMPL, p)
	if err != nil {
		l.Error(err)
	}
}

type SeasonsShowPage struct {
	Season *model.Season
	Flash  *Flash
}

func (p *SeasonsShowPage) Render(w io.Writer) {
	page := pages["seasons_show"]

	err := page.ExecuteTemplate(w, PAGE_TMPL, p)
	if err != nil {
		l.Error(err)
	}
}

type SeasonsNewPage struct {
	Season *model.Season
	Flash  *Flash
}

func (p *SeasonsNewPage) Render(w io.Writer) {
	page := pages["seasons_new"]

	if p.Season == nil {
		p.Season = &model.Season{}
	}

	err := page.ExecuteTemplate(w, PAGE_TMPL, p)
	if err != nil {
		l.Error(err)
	}
}

type SeasonsEditPage struct {
	Season *model.Season
	Flash  *Flash
}

func (p *SeasonsEditPage) Render(w io.Writer) {
	page := pages["seasons_edit"]

	err := page.ExecuteTemplate(w, PAGE_TMPL, p)
	if err != nil {
		l.Error(err)
	}
}
