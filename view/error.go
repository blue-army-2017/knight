package view

import (
	"io"
)

type ErrorPage struct {
	Error error
	Flash *Flash
}

func (p *ErrorPage) Render(w io.Writer) {
	page := pages["error"]

	err := page.ExecuteTemplate(w, "page", p)
	if err != nil {
		l.Error(err)
	}
}

func ShowErrorPage(w io.Writer, err error) {
	page := ErrorPage{
		Error: err,
	}
	page.Render(w)
}
