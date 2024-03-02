package controller

import (
	"net/http"

	"github.com/blue-army-2017/knight/util"
)

var routes = map[string]func(w http.ResponseWriter, r *http.Request){
	"/health":                              getHealth,
	"/":                                    getHome,
	"/members":                             getMembers,
	"/members/new":                         newMember,
	"POST /members/new":                    postNewMember,
	"/members/{id}/edit":                   editMember,
	"POST /members/{id}/edit":              postEditMember,
	"/members/{id}/delete":                 deleteMember,
	"/seasons":                             getSeasons,
	"/seasons/{id}":                        getSeason,
	"/seasons/new":                         newSeason,
	"POST /seasons/new":                    postNewSeason,
	"/seasons/{id}/edit":                   editSeason,
	"POST /seasons/{id}/edit":              postEditSeason,
	"/seasons/{id}/delete":                 deleteSeason,
	"/seasons/{s_id}/games":                getSeasonGames,
	"/seasons/{s_id}/games/new":            newSeasonGame,
	"POST /seasons/{s_id}/games/new":       postNewSeasonGame,
	"/seasons/{s_id}/games/{id}/edit":      editSeasonGame,
	"POST /seasons/{s_id}/games/{id}/edit": postEditSeasonGame,
	"/seasons/{s_id}/games/{id}/delete":    deleteSeasonGame,
}

func GetRoutesMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/static/", assetsHandler())

	for route, handler := range routes {
		if util.Config.Environment == "development" {
			handleFunc := logMiddleware(handler)
			mux.HandleFunc(route, handleFunc)
		} else {
			mux.HandleFunc(route, handler)
		}
	}

	return mux
}

func assetsHandler() http.Handler {
	dir := http.Dir("assets/")
	fs := http.FileServer(dir)
	return http.StripPrefix("/static/", fs)
}
