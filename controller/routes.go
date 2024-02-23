package controller

import "net/http"

func GetRoutesMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", getHealth)

	return mux
}
