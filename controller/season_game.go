package controller

import (
	"net/http"

	"github.com/blue-army-2017/knight/model"
	"github.com/blue-army-2017/knight/view"
)

func getSeasonGames(w http.ResponseWriter, r *http.Request) {
	seasonId := r.PathValue("s_id")

	games, err := model.FindAllSeasonGames(seasonId)
	if err != nil {
		view.ShowErrorPage(w, err)
		return
	}

	page := view.SeasonGamesPage{
		SeasonID: seasonId,
		Games:    games,
	}
	page.Render(w)
}
