package main

import (
	"github.com/blue-army-2017/knight/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("view/**/*")
	routes.Register(router)

	router.Run()
}
