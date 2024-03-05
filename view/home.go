package view

import (
	"html/template"
)

type HomePage struct {
	Flash *Flash
}

func (p *HomePage) Template() *template.Template {
	return pages["home"]
}
