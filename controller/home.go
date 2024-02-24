package controller

import (
	"net/http"

	"github.com/blue-army-2017/knight/view"
)

func getHome(w http.ResponseWriter, r *http.Request) {
	page := view.HomePage{}
	page.Render(w)
}
