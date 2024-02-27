package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/blue-army-2017/knight/model"
	"github.com/blue-army-2017/knight/view"
)

func getSeasons(w http.ResponseWriter, r *http.Request) {
	seasons, err := model.FindAllSeasons()
	if err != nil {
		view.ShowErrorPage(w, err)
		return
	}

	page := view.SeasonsPage{
		Seasons: seasons,
		Flash:   getFlash(w, r),
	}
	page.Render(w)
}

func newSeason(w http.ResponseWriter, r *http.Request) {
	page := view.SeasonsNewPage{}
	page.Render(w)
}

func postNewSeason(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		view.ShowErrorPage(w, err)
		return
	}

	season := parseSeason(r)

	if err := season.Create(); err != nil {
		page := view.SeasonsNewPage{
			Season: &season,
			Flash: &view.Flash{
				Type:    "error",
				Message: err.Error(),
			},
		}
		page.Render(w)
		return
	}

	setFlash(w, "success", fmt.Sprintf("Successfully created season %s", season.Name))
	http.Redirect(w, r, "/seasons", 302)
}

func parseSeason(r *http.Request) model.Season {
	return model.Season{
		Name: strings.TrimSpace(r.FormValue("name")),
	}
}
