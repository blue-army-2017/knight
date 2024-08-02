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
