package controller

import (
	"github.com/blue-army-2017/knight/model"
	"github.com/gin-gonic/gin"
)

type MemberController interface {
	Show() (*Page, error)
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
