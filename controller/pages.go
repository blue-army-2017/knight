package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Page interface {
	Render(ctx *gin.Context)
}

type HtmlPage struct {
	Template string
	Data     any
}

func (p *HtmlPage) Render(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, p.Template, p.Data)
}

type RedirectPage struct {
	Redirect string
}

func (p *RedirectPage) Render(ctx *gin.Context) {
	ctx.Redirect(http.StatusFound, p.Redirect)
}

type ErrorPage struct {
	Error error
}

func (p *ErrorPage) Render(ctx *gin.Context) {
	if errors.Is(p.Error, gorm.ErrRecordNotFound) {
		ctx.AbortWithError(http.StatusBadRequest, p.Error)
	} else {
		ctx.AbortWithError(http.StatusInternalServerError, p.Error)
	}
}
