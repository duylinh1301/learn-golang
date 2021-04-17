package helpers

import (
	"blog/models"
)

var user *models.User

// AuthUser
func AuthUser() *models.User {
	return user
}

// SetUser set user after query
func SetUser(userData *models.User) {
	if user == nil || user.ID != userData.ID {
		user = userData
	}

	return
}
