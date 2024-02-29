package view

import (
	"html/template"
	"io"
)

const PAGE_TMPL = "page"

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
	pages["seasons_show"] = template.Must(template.ParseFiles("templates/seasons/show.html", "templates/layout.html"))
	pages["seasons_new"] = template.Must(template.ParseFiles("templates/seasons/form.html", "templates/seasons/new.html", "templates/layout.html"))
	pages["seasons_edit"] = template.Must(template.ParseFiles("templates/seasons/form.html", "templates/seasons/edit.html", "templates/layout.html"))

	pages["season_games"] = template.Must(template.ParseFiles("templates/seasons/games/index.html", "templates/layout.html"))
	pages["season_games_new"] = template.Must(template.ParseFiles("templates/seasons/games/form.html", "templates/seasons/games/new.html", "templates/layout.html"))
	pages["season_games_edit"] = template.Must(template.ParseFiles("templates/seasons/games/form.html", "templates/seasons/games/edit.html", "templates/layout.html"))
}
