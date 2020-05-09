package main

import (
	"github.com/gin-gonic/gin"
	"practice-golang-gin-webapp/core"
	"practice-golang-gin-webapp/middlewares"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.Use(middlewares.Provide())

	core.SetupRouters(router)

	_ = router.Run(":8080")
}
