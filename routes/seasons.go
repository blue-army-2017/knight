package routes

import (
	"github.com/blue-army-2017/knight/controller"
	"github.com/gin-gonic/gin"
)

var seasonController = controller.NewSeasonController()

func handleSeasons(ctx *gin.Context) {
	page := seasonController.GetIndex()
	page.Render(ctx)
}
