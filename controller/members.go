package controller

import (
	"context"

	"github.com/blue-army-2017/knight/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MemberDto struct {
	ID        string `form:"id"`
	FirstName string `form:"first_name"`
	LastName  string `form:"last_name"`
	Active    bool   `form:"active"`
}

func CreateMemberDto(member repository.Member) MemberDto {
	return MemberDto{
		ID:        member.ID,
		FirstName: member.FirstName,
		LastName:  member.LastName,
		Active:    member.Active > 0.0,
	}
}

func (dto *MemberDto) ToModel() repository.SaveMemberParams {
	var active float64
	if dto.Active {
		active = 1.0
	} else {
		active = 0.0
	}

	return repository.SaveMemberParams{
		ID:        dto.ID,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Active:    active,
	}
}

type MemberController interface {
	GetIndex() Page
	GetNew() Page
	PostNew(member *MemberDto) Page
	GetEdit(id string) Page
	PostEdit(member *MemberDto) Page
}

type DefaultMemberController struct {
	repository repository.Querier
	ctx        context.Context
}

func NewMemberController() MemberController {
	return &DefaultMemberController{
		repository: repository.New(db),
		ctx:        context.Background(),
	}
}

func (c *DefaultMemberController) GetIndex() Page {
	members, err := c.repository.FindAllMembers(c.ctx)
	if err != nil {
		return &ErrorPage{
			Error: err,
		}
	}

	var dtos []MemberDto
	for _, member := range members {
		dto := CreateMemberDto(member)
		dtos = append(dtos, dto)
	}

	return &HtmlPage{
		Template: "pages/members",
		Data: gin.H{
			"Members": dtos,
		},
	}
}

func (c *DefaultMemberController) GetNew() Page {
	member := MemberDto{
		ID:     uuid.NewString(),
		Active: true,
	}

	return &HtmlPage{
		Template: "pages/members/new",
		Data: gin.H{
			"Member": &member,
		},
	}
}

func (c *DefaultMemberController) PostNew(member *MemberDto) Page {
	data := member.ToModel()
	err := c.repository.SaveMember(c.ctx, data)
	if err != nil {
		return &ErrorPage{
			Error: err,
		}
	}

	return &RedirectPage{
		Redirect: "/members",
	}
}

func (c *DefaultMemberController) GetEdit(id string) Page {
	member, err := c.repository.FindMemberById(c.ctx, id)
	if err != nil {
		return &ErrorPage{
			Error: err,
		}
	}

	return &HtmlPage{
		Template: "pages/members/edit",
		Data: gin.H{
			"Member": CreateMemberDto(member),
		},
	}
}

func (c *DefaultMemberController) PostEdit(member *MemberDto) Page {
	data := member.ToModel()
	if err := c.repository.SaveMember(c.ctx, data); err != nil {
		return &ErrorPage{
			Error: err,
		}
	}

	return &RedirectPage{
		Redirect: "/members",
	}
}
