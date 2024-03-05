package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/blue-army-2017/knight/model"
	"github.com/blue-army-2017/knight/view"
)

func getMembers(w http.ResponseWriter, r *http.Request) {
	members, err := model.FindAllMembers()
	if err != nil {
		view.ShowErrorPage(w, err)
		return
	}

	page := view.MembersPage{
		Members: members,
		Flash:   getFlash(w, r),
	}
	view.RenderPage(w, &page)
}

func newMember(w http.ResponseWriter, r *http.Request) {
	page := view.MembersNewPage{
		Member: &model.Member{
			Active: true,
		},
	}
	view.RenderPage(w, &page)
}

func postNewMember(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		view.ShowErrorPage(w, err)
		return
	}

	var member model.Member
	parseMember(r, &member)

	if err := member.Create(); err != nil {
		page := view.MembersNewPage{
			Member: &member,
			Flash: &view.Flash{
				Type:    "error",
				Message: err.Error(),
			},
		}
		view.RenderPage(w, &page)
		return
	}

	setFlash(w, "success", fmt.Sprintf("Successfully created member %s %s", member.FirstName, member.LastName))
	http.Redirect(w, r, "/members", 302)
}

func editMember(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	member, err := model.FindMemberByID(id)
	if err != nil {
		view.ShowErrorPage(w, err)
		return
	}

	page := view.MembersEditPage{
		Member: &member,
	}
	view.RenderPage(w, &page)
}

func postEditMember(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := r.ParseForm(); err != nil {
		view.ShowErrorPage(w, err)
		return
	}

	member, err := model.FindMemberByID(id)
	if err != nil {
		view.ShowErrorPage(w, err)
		return
	}

	parseMember(r, &member)

	if err := member.Update(); err != nil {
		page := view.MembersEditPage{
			Member: &member,
			Flash: &view.Flash{
				Type:    "error",
				Message: err.Error(),
			},
		}
		view.RenderPage(w, &page)
		return
	}

	setFlash(w, "success", fmt.Sprintf("Successfully updated member %s %s", member.FirstName, member.LastName))
	http.Redirect(w, r, "/members", 302)
}

func deleteMember(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	member, err := model.FindMemberByID(id)
	if err != nil {
		view.ShowErrorPage(w, err)
		return
	}

	if err := member.Delete(); err != nil {
		view.ShowErrorPage(w, err)
		return
	}

	setFlash(w, "success", fmt.Sprintf("Successfully deleted member %s %s", member.FirstName, member.LastName))
	http.Redirect(w, r, "/members", 302)
}

func parseMember(r *http.Request, member *model.Member) {
	if member == nil {
		return
	}

	member.FirstName = strings.TrimSpace(r.FormValue("first_name"))
	member.LastName = strings.TrimSpace(r.FormValue("last_name"))
	member.Active = r.FormValue("active") == "on"
}
