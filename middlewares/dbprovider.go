package middlewares

import (
	"github.com/gin-gonic/gin"
	"practice-golang-gin-webapp/core"
)

func Provide() gin.HandlerFunc {
	db := core.SetupDB()

	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
