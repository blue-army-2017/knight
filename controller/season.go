package controller

import (
	"net/http"

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
