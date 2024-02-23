package controller

import (
	"net/http"

	"github.com/blue-army-2017/knight/util"
	"github.com/blue-army-2017/knight/view"
)

func getHome(w http.ResponseWriter, r *http.Request) {
	util.LogDebug("request received", "path", "/", "method", "GET")

	page := view.HomePage{}
	page.Render(w)
}
