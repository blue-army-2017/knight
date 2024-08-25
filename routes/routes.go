package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	// Static assets
	router.Static("/static", "./assets")
	// Monitoring endpoints
	router.GET("/health", handleHealth)
	// Homepage
	router.GET("/", handleIndex)

	// Members Module
	router.GET("/members", handleMembers)
	router.GET("/members/new", handleMembersNew)
	router.POST("/members/new", handleMembersNewPost)
	router.GET("/members/:id", handleMembersEdit)
	router.POST("/members/:id", handleMembersEditPost)

	// Seasons Module
	router.GET("/seasons", handleSeasons)
	router.GET("/seasons/new", handleSeasonsNew)
	router.POST("/seasons/new", handleSeasonsNewPost)
	router.GET("/seasons/:id", handleSeasonsEdit)
	router.POST("/seasons/:id", handleSeasonsEditPost)

	// Season Games Module
	router.GET("/games", handleSeasonGames)
	router.GET("/games/new", handleSeasonGamesNew)
	router.POST("/games/new", handleSeasonGamesNewPost)
	router.GET("/games/:id", handleSeasonGamesEdit)
	router.POST("/games/:id", handleSeasonGamesEditPost)

	// Presence Module
	router.GET("/presence", handlePresence)
}

func handleHealth(ctx *gin.Context) {
	ctx.String(http.StatusOK, "UP")
}

func handleIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "pages/index", nil)
}
