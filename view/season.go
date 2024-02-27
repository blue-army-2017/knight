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

	err := page.ExecuteTemplate(w, "page", p)
	if err != nil {
		l.Error(err)
	}
}
