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
}

func handleHealth(ctx *gin.Context) {
	ctx.String(http.StatusOK, "UP")
}

func handleIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "pages/index", nil)
}
