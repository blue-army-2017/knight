package view

import (
	"io"

	"github.com/blue-army-2017/knight/util"
)

type ErrorPage struct {
	Error error
	Flash *Flash
}

func (p *ErrorPage) Render(w io.Writer) {
	page := pages["error"]

	err := page.ExecuteTemplate(w, "page", p)
	if err != nil {
		util.LogError(err.Error())
	}
}

func ShowErrorPage(w io.Writer, err error) {
	page := ErrorPage{
		Error: err,
	}
	page.Render(w)
}
