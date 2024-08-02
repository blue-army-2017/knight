package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Page struct {
	Template string
	Data     any
}

func (p *Page) Render(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, p.Template, p.Data)
}
