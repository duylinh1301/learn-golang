package implement

import (
	"blog/models"
	"fmt"
)

type GoogleAuth struct {
}

func NewGoogleAuth() *GoogleAuth {
	return &GoogleAuth{}
}

func (*GoogleAuth) Login(map[string]interface{}) (interface{}, *models.User) {
	fmt.Println("google login")

	return true, &models.User{}
}
