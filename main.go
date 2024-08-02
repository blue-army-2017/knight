package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("view/**/*")
	RegisterRoutes(router)

	router.Run()
}
