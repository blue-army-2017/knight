package controller

import (
	"fmt"
	"net/http"

	"github.com/blue-army-2017/knight/model"
	"github.com/blue-army-2017/knight/view"
)

func getSeasonGames(w http.ResponseWriter, r *http.Request) {
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

	page := view.SeasonGamesPage{
		Season: &season,
		Games:  games,
		Flash:  getFlash(w, r),
	}
	page.Render(w)
}

func newSeasonGame(w http.ResponseWriter, r *http.Request) {
	seasonId := r.PathValue("s_id")

	season, err := model.FindSeasonByID(seasonId)
	if err != nil {
		view.ShowErrorPage(w, err)
		return
	}

	members, err := model.FindAllMembers()
	if err != nil {
		view.ShowErrorPage(w, err)
		return
	}

	page := view.SeasonGamesNewPage{
		Season:  &season,
		Members: members,
	}
	page.Render(w)
}

func postNewSeasonGame(w http.ResponseWriter, r *http.Request) {
	seasonId := r.PathValue("s_id")
	if err := r.ParseForm(); err != nil {
		view.ShowErrorPage(w, err)
		return
	}

	season, err := model.FindSeasonByID(seasonId)
	if err != nil {
		view.ShowErrorPage(w, err)
		return
	}

	members, err := model.FindAllMembers()
	if err != nil {
		view.ShowErrorPage(w, err)
		return
	}

	game := model.SeasonGame{
		SeasonID: seasonId,
	}
	parseSeasonGame(r, &game, members)

	if err := game.Create(); err != nil {
		page := view.SeasonGamesNewPage{
			Season:  &season,
			Game:    &game,
			Members: members,
			Flash: &view.Flash{
				Type:    "error",
				Message: err.Error(),
			},
		}
		page.Render(w)
		return
	}

	setFlash(w, "success", fmt.Sprintf("Successfully created game %s (%s)", game.Opponent, game.Date))
	http.Redirect(w, r, fmt.Sprintf("/seasons/%s/games", seasonId), 302)
}

func editSeasonGame(w http.ResponseWriter, r *http.Request) {
	seasonId := r.PathValue("s_id")
	id := r.PathValue("id")

	season, err := model.FindSeasonByID(seasonId)
	if err != nil {
		view.ShowErrorPage(w, err)
		return
	}

	game, err := model.FindSeasonGameByID(id)
	if err != nil {
		view.ShowErrorPage(w, err)
		return
	}

	members, err := model.FindAllMembers()
	if err != nil {
		view.ShowErrorPage(w, err)
		return
	}

	page := view.SeasonGamesEditPage{
		Season:  &season,
		Game:    &game,
		Members: members,
	}
	page.Render(w)
}

func postEditSeasonGame(w http.ResponseWriter, r *http.Request) {
	seasonId := r.PathValue("s_id")
	id := r.PathValue("id")
	if err := r.ParseForm(); err != nil {
		view.ShowErrorPage(w, err)
		return
	}

	season, err := model.FindSeasonByID(seasonId)
	if err != nil {
		view.ShowErrorPage(w, err)
		return
	}

	game, err := model.FindSeasonGameByID(id)
	if err != nil {
		view.ShowErrorPage(w, err)
		return
	}

	members, err := model.FindAllMembers()
	if err != nil {
		view.ShowErrorPage(w, err)
		return
	}

	parseSeasonGame(r, &game, members)

	if err := game.Update(); err != nil {
		page := view.SeasonGamesEditPage{
			Season:  &season,
			Game:    &game,
			Members: members,
			Flash: &view.Flash{
				Type:    "error",
				Message: err.Error(),
			},
		}
		page.Render(w)
		return
	}

	setFlash(w, "success", fmt.Sprintf("Successfully updated game %s (%s)", game.Opponent, game.Date))
	http.Redirect(w, r, fmt.Sprintf("/seasons/%s/games", seasonId), 302)
}

func deleteSeasonGame(w http.ResponseWriter, r *http.Request) {
	seasonId := r.PathValue("s_id")
	id := r.PathValue("id")

	game, err := model.FindSeasonGameByID(id)
	if err != nil {
		view.ShowErrorPage(w, err)
		return
	}

	if err := game.Delete(); err != nil {
		view.ShowErrorPage(w, err)
		return
	}

	setFlash(w, "success", fmt.Sprintf("Successfully deleted game %s (%s)", game.Opponent, game.Date))
	http.Redirect(w, r, fmt.Sprintf("/seasons/%s/games", seasonId), 302)
}

func parseSeasonGame(r *http.Request, game *model.SeasonGame, members []model.Member) {
	if game == nil {
		return
	}

	game.Opponent = r.FormValue("opponent")
	game.Home = r.FormValue("home") == "on"
	game.Date = r.FormValue("date")
	game.Mode = r.FormValue("mode")

	presentMembers := []model.Member{}
	for _, member := range members {
		if r.FormValue(member.ID) == "on" {
			presentMembers = append(presentMembers, member)
		}
	}
	game.PresentMembers = presentMembers
}
