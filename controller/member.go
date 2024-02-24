package controller

import (
	"fmt"
	"net/http"
	"strings"

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

func newMember(w http.ResponseWriter, r *http.Request) {
	page := view.MembersNewPage{}
	page.Render(w)
}

func postNewMember(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		util.LogError(err.Error())
		http.Error(w, err.Error(), 400)
		return
	}

	member := model.Member{
		FirstName: strings.TrimSpace(r.FormValue("first_name")),
		LastName:  strings.TrimSpace(r.FormValue("last_name")),
		Active:    r.FormValue("active") == "on",
	}

	if err := member.Create(); err != nil {
		page := view.MembersNewPage{
			Member: &member,
			Flash: &view.Flash{
				Type:    "error",
				Message: err.Error(),
			},
		}
		page.Render(w)
		return
	}

	setFlash(w, "success", fmt.Sprintf("Successfully created member %s %s", member.FirstName, member.LastName))
	http.Redirect(w, r, "/members", 302)
}

func editMember(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	member, err := model.FindMemberByID(id)
	if err != nil {
		util.LogError(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	page := view.MembersEditPage{
		Member: &member,
	}
	page.Render(w)
}

func postEditMember(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := r.ParseForm(); err != nil {
		util.LogError(err.Error())
		http.Error(w, err.Error(), 400)
		return
	}

	member := model.Member{
		ID:        id,
		FirstName: strings.TrimSpace(r.FormValue("first_name")),
		LastName:  strings.TrimSpace(r.FormValue("last_name")),
		Active:    r.FormValue("active") == "on",
	}

	if err := member.Update(); err != nil {
		page := view.MembersEditPage{
			Member: &member,
			Flash: &view.Flash{
				Type:    "error",
				Message: err.Error(),
			},
		}
		page.Render(w)
		return
	}

	setFlash(w, "success", fmt.Sprintf("Successfully updated member %s %s", member.FirstName, member.LastName))
	http.Redirect(w, r, "/members", 302)
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
