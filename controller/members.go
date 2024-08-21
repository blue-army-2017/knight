package controller

import (
	"github.com/blue-army-2017/knight/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MemberDto struct {
	ID        string `form:"id"`
	FirstName string `form:"first_name"`
	LastName  string `form:"last_name"`
	Active    bool   `form:"active"`
}

func CreateMemberDto(member *model.Member) *MemberDto {
	return &MemberDto{
		ID:        member.ID,
		FirstName: member.FirstName,
		LastName:  member.LastName,
		Active:    member.Active,
	}
}

func (dto *MemberDto) ToModel() *model.Member {
	return &model.Member{
		ID:        dto.ID,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Active:    dto.Active,
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
	repository model.CRUDRepository[model.Member]
}

func NewMemberController() MemberController {
	return &DefaultMemberController{
		repository: model.NewCRUDRepository[model.Member](),
	}
}

func (c *DefaultMemberController) GetIndex() Page {
	members, err := c.repository.FindAll("last_name", "first_name")
	if err != nil {
		return &ErrorPage{
			Error: err,
		}
	}

	var dtos []MemberDto
	for _, member := range members {
		dto := CreateMemberDto(&member)
		dtos = append(dtos, *dto)
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
	err := c.repository.Create(data)
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
	member, err := c.repository.FindById(id)
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
	if err := c.repository.Update(data); err != nil {
		return &ErrorPage{
			Error: err,
		}
	}

	return &RedirectPage{
		Redirect: "/members",
	}
}
