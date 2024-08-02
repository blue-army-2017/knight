package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Page struct {
	Template string
	Data     any
	Redirect string
	Error    error
}

func (p *Page) Render(ctx *gin.Context) {
	err := p.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	} else if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if p.Redirect != "" {
		ctx.Redirect(http.StatusFound, p.Redirect)
		return
	}

	ctx.HTML(http.StatusOK, p.Template, p.Data)
}
