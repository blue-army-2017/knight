package main

import (
	"net/http"

	"github.com/blue-army-2017/knight/controller"
	"github.com/gin-gonic/gin"
)

var (
	memberController = controller.NewMemberController()
)

func RegisterRoutes(router *gin.Engine) {
	// Static assets
	router.Static("/static", "./assets")
	// Monitoring endpoints
	router.GET("/health", handleHealth)

	// Members Module
	router.GET("/members", handleMembers)
}

func handleHealth(ctx *gin.Context) {
	ctx.String(http.StatusOK, "UP")
}

func handleMembers(ctx *gin.Context) {
	page, err := memberController.Show()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	page.Render(ctx)
}
