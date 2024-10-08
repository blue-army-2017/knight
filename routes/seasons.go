package routes

import (
	"net/http"

	"github.com/blue-army-2017/knight/controller"
	"github.com/gin-gonic/gin"
)

var seasonController = controller.NewSeasonController()

func handleSeasons(ctx *gin.Context) {
	page := seasonController.GetIndex()
	page.Render(ctx)
}

func handleSeasonsNew(ctx *gin.Context) {
	page := seasonController.GetNew()
	page.Render(ctx)
}

func handleSeasonsNewPost(ctx *gin.Context) {
	var season controller.SeasonDto
	if err := ctx.Bind(&season); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	page := seasonController.PostNew(&season)
	page.Render(ctx)
}

func handleSeasonsEdit(ctx *gin.Context) {
	id := ctx.Param("id")
	page := seasonController.GetEdit(id)
	page.Render(ctx)
}

func handleSeasonsEditPost(ctx *gin.Context) {
	var season controller.SeasonDto
	if err := ctx.Bind(&season); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	_, isDelete := ctx.GetQuery("delete")

	page := seasonController.PostEdit(&season, isDelete)
	page.Render(ctx)
}
