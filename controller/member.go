package controller

import (
	"fmt"
	"net/http"

	"github.com/blue-army-2017/knight/model"
	"github.com/blue-army-2017/knight/util"
	"github.com/blue-army-2017/knight/view"
)

func getMembers(w http.ResponseWriter, r *http.Request) {
	members, err := model.FindAllMembers()
	if err != nil {
		util.LogError(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	page := view.MembersPage{
		Members: members,
		Flash:   getFlash(w, r),
	}
	page.Render(w)
}

func deleteMember(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	member, err := model.FindMemberByID(id)
	if err != nil {
		util.LogError(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	if err := member.Delete(); err != nil {
		util.LogError(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	setFlash(w, "success", fmt.Sprintf("Successfully deleted member %s %s", member.FirstName, member.LastName))
	http.Redirect(w, r, "/members", 302)
}
