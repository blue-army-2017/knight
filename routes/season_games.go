package routes

import (
	"net/http"

	"github.com/blue-army-2017/knight/controller"
	"github.com/gin-gonic/gin"
)

var seasonGamesController = controller.NewSeasonGameController()

func handleSeasonGames(ctx *gin.Context) {
	seasonId := ctx.Query("season")
	page := seasonGamesController.GetIndex(seasonId)
	page.Render(ctx)
}

func handleSeasonGamesNew(ctx *gin.Context) {
	page := seasonGamesController.GetNew()
	page.Render(ctx)
}

func handleSeasonGamesNewPost(ctx *gin.Context) {
	var game controller.SeasonGameDto
	if err := ctx.Bind(&game); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	page := seasonGamesController.PostNew(&game)
	page.Render(ctx)
}
