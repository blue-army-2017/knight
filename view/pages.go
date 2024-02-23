package view

import (
	"html/template"
	"io"
)

type Page interface {
	Render(w io.Writer)
}

var pages map[string]*template.Template

func init() {
	pages = make(map[string]*template.Template)

	pages["home"] = template.Must(template.ParseFiles("templates/home.html", "templates/layout.html"))
}
