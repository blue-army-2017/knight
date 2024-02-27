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
	pages["error"] = template.Must(template.ParseFiles("templates/error.html", "templates/layout.html"))

	pages["members"] = template.Must(template.ParseFiles("templates/members/index.html", "templates/layout.html"))
	pages["members_new"] = template.Must(template.ParseFiles("templates/members/form.html", "templates/members/new.html", "templates/layout.html"))
	pages["members_edit"] = template.Must(template.ParseFiles("templates/members/form.html", "templates/members/edit.html", "templates/layout.html"))

	pages["seasons"] = template.Must(template.ParseFiles("templates/seasons/index.html", "templates/layout.html"))
}
