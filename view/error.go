package view

import (
	"html/template"
	"io"
)

type ErrorPage struct {
	Error error
	Flash *Flash
}

func (p *ErrorPage) Template() *template.Template {
	return pages["error"]
}

func ShowErrorPage(w io.Writer, err error) {
	page := ErrorPage{
		Error: err,
	}
	RenderPage(w, &page)
}
