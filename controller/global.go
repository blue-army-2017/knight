package controller

import (
	"github.com/gin-gonic/gin"
)

type Page struct {
	StatusCode int
	Template   string
	Data       any
}

func (p *Page) Render(ctx *gin.Context) {
	ctx.HTML(p.StatusCode, p.Template, p.Data)
}
