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

	err := page.ExecuteTemplate(w, "page", p)
	if err != nil {
		l.Error(err)
	}
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

	err := page.ExecuteTemplate(w, "page", p)
	if err != nil {
		l.Error(err)
	}
}

type MembersEditPage struct {
	Member *model.Member
	Flash  *Flash
}

func (p *MembersEditPage) Render(w io.Writer) {
	page := pages["members_edit"]

	err := page.ExecuteTemplate(w, "page", p)
	if err != nil {
		l.Error(err)
	}
}
