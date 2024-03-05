package view

import (
	"html/template"

	"github.com/blue-army-2017/knight/model"
)

type MembersPage struct {
	Members []model.Member
	Flash   *Flash
}

func (p *MembersPage) Template() *template.Template {
	return pages["members"]
}

type MembersNewPage struct {
	Member *model.Member
	Flash  *Flash
}

func (p *MembersNewPage) Template() *template.Template {
	return pages["members_new"]
}

type MembersEditPage struct {
	Member *model.Member
	Flash  *Flash
}

func (p *MembersEditPage) Template() *template.Template {
	return pages["members_edit"]
}
