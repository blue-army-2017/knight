package controller

import "net/http"

var routes = map[string]func(w http.ResponseWriter, r *http.Request){
	"/health":                 getHealth,
	"/":                       getHome,
	"/members":                getMembers,
	"/members/new":            newMember,
	"POST /members/new":       postNewMember,
	"/members/{id}/edit":      editMember,
	"POST /members/{id}/edit": postEditMember,
	"/members/{id}/delete":    deleteMember,
	"/seasons":                getSeasons,
	"/seasons/new":            newSeason,
	"POST /seasons/new":       postNewSeason,
	"/seasons/{id}/edit":      editSeason,
	"POST /seasons/{id}/edit": postEditSeason,
}

func GetRoutesMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/static/", assetsHandler())

	for route, handler := range routes {
		handleFunc := logMiddleware(handler)
		mux.HandleFunc(route, handleFunc)
	}

	return mux
}

func assetsHandler() http.Handler {
	dir := http.Dir("assets/")
	fs := http.FileServer(dir)
	return http.StripPrefix("/static/", fs)
}
