package auth

import "fmt"

type JWTAuth struct {
}

func NewJWTAuth() *JWTAuth {
	return &JWTAuth{}
}

func (*JWTAuth) Login() {
	fmt.Println("traditional login")
}
