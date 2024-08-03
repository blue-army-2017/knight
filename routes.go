package main

import (
	"net/http"

	"github.com/blue-army-2017/knight/controller"
	"github.com/blue-army-2017/knight/model"
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

func handleMembers(ctx *gin.Context) {
	page := memberController.Show()
	page.Render(ctx)
}

func handleMembersNew(ctx *gin.Context) {
	page := memberController.New()
	page.Render(ctx)
}

func handleMembersNewPost(ctx *gin.Context) {
	var member model.Member
	if err := ctx.Bind(&member); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	page := memberController.PostNew(&member)
	page.Render(ctx)
}

func handleMembersEdit(ctx *gin.Context) {
	id := ctx.Param("id")
	page := memberController.Edit(id)
	page.Render(ctx)
}

func handleMembersEditPost(ctx *gin.Context) {
	var member model.Member
	if err := ctx.Bind(&member); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	page := memberController.PostEdit(&member)
	page.Render(ctx)
}
