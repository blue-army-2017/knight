package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// Static assets
	router.Static("/static", "./assets")
	// Monitoring endpoints
	router.GET("/health", handleHealth)
}

func handleHealth(ctx *gin.Context) {
	ctx.String(http.StatusOK, "UP")
}
