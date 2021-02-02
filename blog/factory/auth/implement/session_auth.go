package implement

import (
	"blog/models"
	"fmt"
)

type SessionAuth struct{}

func NewSessionAuth() *GoogleAuth {
	return &GoogleAuth{}
}

func (*SessionAuth) Login(map[string]interface{}) (interface{}, *models.User) {
	fmt.Println("Session login")

	return false, nil
}
