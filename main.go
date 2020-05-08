package main

import (
	"github.com/gin-gonic/gin"
	"practice-golang-gin-webapp/core"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	_ = core.SetupDb()

	core.SetupRouters(router)

	_ = router.Run(":8080")
}
