package controller

import "net/http"

func GetRoutesMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/static/", assetsHandler())
	mux.HandleFunc("/health", getHealth)

	mux.HandleFunc("/", getHome)

	mux.HandleFunc("/members", getMembers)
	mux.HandleFunc("/members/{id}/delete", deleteMember)

	return mux
}

func assetsHandler() http.Handler {
	dir := http.Dir("assets/")
	fs := http.FileServer(dir)
	return http.StripPrefix("/static/", fs)
}
