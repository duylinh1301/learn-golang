package auth

import (
	"blog/factory/auth/implement"
	"blog/factory/auth/interfaces"
)

const TypeJWT = 1
const TypeSession = 2
const TypeGoogle = 3
const TypeFacebook = 4

type Auth struct {
	AuthMethod interfaces.LoginInterface
}

func NewAuthFactory(authType int) *Auth {

	switch authType {

	case TypeJWT:
		return &Auth{
			AuthMethod: implement.NewJWTAuth(),
		}

	case TypeGoogle:
		return &Auth{
			AuthMethod: implement.NewGoogleAuth(),
		}

	default:
		return &Auth{
			AuthMethod: implement.NewSessionAuth(),
		}
	}
}
