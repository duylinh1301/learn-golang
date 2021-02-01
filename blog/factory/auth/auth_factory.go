package auth

import (
	"blog/config"
	"blog/factory/auth/interfaces"
	"fmt"
)

type AuthFactory struct {
	AuthMethod interfaces.LoginInterface
}

func NewAuthFactory() *AuthFactory {

	authConfig := config.Auth

	fmt.Println(authConfig["auth_type"]["jwt"])

	return &AuthFactory{
		AuthMethod: &GoogleAuth{},
	}
}

func (*AuthFactory) getAuthMethod(loginMethod int) interfaces.LoginInterface {

	// switch loginMethod {

	// case authConfig["auth_type"]["jwt"]:

	// }

	return &GoogleAuth{}
}
