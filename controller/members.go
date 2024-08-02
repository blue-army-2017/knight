package controller

import (
	"github.com/blue-army-2017/knight/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MemberController interface {
	Show() (*Page, error)
	New() *Page
	NewPost(member *model.Member) error
}

type DefaultMemberController struct {
	repository model.MemberRepository
}

func NewMemberController() MemberController {
	return &DefaultMemberController{
		repository: model.NewMemberRepository(),
	}
}

func (c *DefaultMemberController) Show() (*Page, error) {
	members, err := c.repository.FindAll()
	if err != nil {
		return nil, err
	}

	return &Page{
		Template: "pages/members",
		Data: gin.H{
			"Members": members,
		},
	}, nil
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

func (c *DefaultMemberController) NewPost(member *model.Member) error {
	return c.repository.Create(member)
}
