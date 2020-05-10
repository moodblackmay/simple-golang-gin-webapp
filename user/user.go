package user

import (
	"errors"
	"github.com/jinzhu/gorm"
	"strings"
)

//User model definition
type User struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//Create new user
func RegisterNewUser(username, password string, db *gorm.DB) (*User, error) {
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("the password can't be empty")
	} else if !IsUsernameAvailable(username, db) {
		return nil, errors.New("the username isn't available")
	}

	NewUser := User{Username: username, Password: password}

	return &NewUser, nil
}

//Check if username don't exist
func IsUsernameAvailable(username string, db *gorm.DB) bool {
	user := User{}

	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return true
	}

	return false
}
