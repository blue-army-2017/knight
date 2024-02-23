package view

import (
	"io"

	"github.com/blue-army-2017/knight/util"
)

type HomePage struct{}

func (p *HomePage) Render(w io.Writer) {
	page := pages["home"]

	err := page.ExecuteTemplate(w, "page", nil)
	if err != nil {
		util.LogError(err.Error())
	}
}
