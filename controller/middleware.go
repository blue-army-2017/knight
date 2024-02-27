package controller

import (
	"net/http"
)

func logMiddleware(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		method := r.Method
		l.Debugw("request received", "path", path, "method", method)

		next(w, r)
	}
}
