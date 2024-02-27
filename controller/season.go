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

func getSeason(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	season, err := model.FindSeasonByID(id)
	if err != nil {
		view.ShowErrorPage(w, err)
		return
	}

	page := view.SeasonsShowPage{
		Season: &season,
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

	var season model.Season
	parseSeason(r, &season)

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

func editSeason(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	season, err := model.FindSeasonByID(id)
	if err != nil {
		view.ShowErrorPage(w, err)
		return
	}

	page := view.SeasonsEditPage{
		Season: &season,
	}
	page.Render(w)
}

func postEditSeason(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := r.ParseForm(); err != nil {
		view.ShowErrorPage(w, err)
		return
	}

	season, err := model.FindSeasonByID(id)
	if err != nil {
		view.ShowErrorPage(w, err)
		return
	}

	parseSeason(r, &season)

	if err := season.Update(); err != nil {
		page := view.SeasonsEditPage{
			Season: &season,
			Flash: &view.Flash{
				Type:    "error",
				Message: err.Error(),
			},
		}
		page.Render(w)
		return
	}

	setFlash(w, "success", fmt.Sprintf("Successfully updated season %s", season.Name))
	http.Redirect(w, r, "/seasons", 302)
}

func deleteSeason(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	season, err := model.FindSeasonByID(id)
	if err != nil {
		view.ShowErrorPage(w, err)
		return
	}

	if err := season.Delete(); err != nil {
		view.ShowErrorPage(w, err)
		return
	}

	setFlash(w, "success", fmt.Sprintf("Successfully deleted season %s", season.Name))
	http.Redirect(w, r, "/seasons", 302)
}

func parseSeason(r *http.Request, season *model.Season) {
	if season == nil {
		return
	}

	season.Name = strings.TrimSpace(r.FormValue("name"))
}
