package controller

import (
	"github.com/blue-army-2017/knight/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MemberController interface {
	Show() *Page
	New() *Page
	PostNew(member *model.Member) *Page
	Edit(id string) *Page
	PostEdit(member *model.Member) *Page
}

type DefaultMemberController struct {
	repository model.MemberRepository
}

func NewMemberController() MemberController {
	return &DefaultMemberController{
		repository: model.NewMemberRepository(),
	}
}

func (c *DefaultMemberController) Show() *Page {
	members, err := c.repository.FindAll()
	if err != nil {
		return &Page{
			Error: err,
		}
	}

	return &Page{
		Template: "pages/members",
		Data: gin.H{
			"Members": members,
		},
	}
}

func (c *DefaultMemberController) New() *Page {
	member := model.Member{
		ID:     uuid.New().String(),
		Active: true,
	}

	return &Page{
		Template: "pages/members/new",
		Data: gin.H{
			"Member": member,
		},
	}
}

func (c *DefaultMemberController) PostNew(member *model.Member) *Page {
	err := c.repository.Create(member)
	if err != nil {
		return &Page{
			Error: err,
		}
	}

	return &Page{
		Redirect: "/members",
	}
}

func (c *DefaultMemberController) Edit(id string) *Page {
	member, err := c.repository.FindById(id)
	if err != nil {
		return &Page{
			Error: err,
		}
	}

	return &Page{
		Template: "pages/members/edit",
		Data: gin.H{
			"Member": member,
		},
	}
}

func (c *DefaultMemberController) PostEdit(member *model.Member) *Page {
	if err := c.repository.Update(member); err != nil {
		return &Page{
			Error: err,
		}
	}

	return &Page{
		Redirect: "/members",
	}
}
