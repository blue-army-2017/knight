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

type MembersNewPage struct {
	Member *model.Member
	Flash  *Flash
}

func (p *MembersNewPage) Render(w io.Writer) {
	page := pages["members_new"]

	if p.Member == nil {
		p.Member = &model.Member{
			Active: true,
		}
	}

	page.ExecuteTemplate(w, "page", p)
}

type MembersEditPage struct {
	Member *model.Member
	Flash  *Flash
}

func (p *MembersEditPage) Render(w io.Writer) {
	page := pages["members_edit"]

	page.ExecuteTemplate(w, "page", p)
}
