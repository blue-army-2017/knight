package routes

import (
	"net/http"

	"github.com/blue-army-2017/knight/controller"
	"github.com/blue-army-2017/knight/model"
	"github.com/gin-gonic/gin"
)

var memberController = controller.NewMemberController()

func handleMembers(ctx *gin.Context) {
	page := memberController.Show()
	page.Render(ctx)
}

func handleMembersNew(ctx *gin.Context) {
	page := memberController.New()
	page.Render(ctx)
}

func handleMembersNewPost(ctx *gin.Context) {
	var member model.Member
	if err := ctx.Bind(&member); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	page := memberController.PostNew(&member)
	page.Render(ctx)
}

func handleMembersEdit(ctx *gin.Context) {
	id := ctx.Param("id")
	page := memberController.Edit(id)
	page.Render(ctx)
}

func handleMembersEditPost(ctx *gin.Context) {
	var member model.Member
	if err := ctx.Bind(&member); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	page := memberController.PostEdit(&member)
	page.Render(ctx)
}
