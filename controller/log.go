package controller

import (
	"net/http"

	"github.com/blue-army-2017/knight/util"
)

func logMiddleware(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		method := r.Method
		util.LogDebug("request received", "path", path, "method", method)

		next(w, r)
	}
}
