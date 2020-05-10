package user

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func ShowRegistrationPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func Register(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	username := c.PostForm("username")
	password := c.PostForm("password")

	if user, err := RegisterNewUser(username, password, db); err == nil {
		db.Create(user)
		c.HTML(http.StatusOK, "login-successful.html", nil)
	} else {
		err := err.Error()
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"ErrorTitle": "Registration Failed",
			"Error":      err,
		})
	}
}
