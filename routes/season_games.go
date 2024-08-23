package routes

import (
	"github.com/blue-army-2017/knight/controller"
	"github.com/gin-gonic/gin"
)

var seasonGamesController = controller.NewSeasonGameController()

func handleSeasonGames(ctx *gin.Context) {
	seasonId := ctx.Query("season")
	page := seasonGamesController.GetIndex(seasonId)
	page.Render(ctx)
}
