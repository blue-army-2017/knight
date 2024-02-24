package view

import (
	"io"

	"github.com/blue-army-2017/knight/model"
)

type MembersPage struct {
	Members []model.Member
	Flash   *Flash
}

func (p *MembersPage) Render(w io.Writer) {
	page := pages["members"]

	page.ExecuteTemplate(w, "page", p)
}
