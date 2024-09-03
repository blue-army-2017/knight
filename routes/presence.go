package routes

import (
	"net/http"

	"github.com/blue-army-2017/knight/controller"
	"github.com/gin-gonic/gin"
)

var presenceController = controller.NewPresenceController()

func handlePresence(ctx *gin.Context) {
	page := presenceController.GetIndex()
	page.Render(ctx)
}

func handlePresenceEdit(ctx *gin.Context) {
	gameId := ctx.Param("gameId")
	page := presenceController.GetEdit(gameId)
	page.Render(ctx)
}

func handlePresenceEditPost(ctx *gin.Context) {
	gameId := ctx.Param("gameId")
	var data struct {
		PresentMembers []string `form:"present_members[]"`
	}
	if err := ctx.Bind(&data); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	page := presenceController.GetEditPost(gameId, data.PresentMembers)
	page.Render(ctx)
}
