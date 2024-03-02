package controller

import (
	"net/http"

	"github.com/blue-army-2017/knight/model"
	"github.com/blue-army-2017/knight/view"
)

func getSeasonPresence(w http.ResponseWriter, r *http.Request) {
	seasonId := r.PathValue("s_id")

	season, err := model.FindSeasonByID(seasonId)
	if err != nil {
		view.ShowErrorPage(w, err)
		return
	}

	games, err := model.FindAllSeasonGames(seasonId)
	if err != nil {
		view.ShowErrorPage(w, err)
		return
	}

	members, err := model.FindAllMembers()
	if err != nil {
		view.ShowErrorPage(w, err)
		return
	}

	page := view.SeasonPresencePage{
		Season:  &season,
		Games:   games,
		Members: members,
	}
	page.Render(w)
}
