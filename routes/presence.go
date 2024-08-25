package routes

import (
	"github.com/blue-army-2017/knight/controller"
	"github.com/gin-gonic/gin"
)

var presenceController = controller.NewPresenceController()

func handlePresence(ctx *gin.Context) {
	page := presenceController.GetIndex()
	page.Render(ctx)
}
