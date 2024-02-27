package view

import (
	"io"
)

type HomePage struct{}

func (p *HomePage) Render(w io.Writer) {
	page := pages["home"]

	err := page.ExecuteTemplate(w, PAGE_TMPL, nil)
	if err != nil {
		l.Error(err)
	}
}
