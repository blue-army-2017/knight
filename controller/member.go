package controller

import (
	"net/http"

	"github.com/blue-army-2017/knight/model"
	"github.com/blue-army-2017/knight/util"
	"github.com/blue-army-2017/knight/view"
)

func getMembers(w http.ResponseWriter, r *http.Request) {
	util.LogDebug("request received", "path", "/members", "method", "GET")

	members, err := model.FindAllMembers()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	page := view.MembersPage{
		Members: members,
	}
	page.Render(w)
}
