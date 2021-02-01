package auth

import "fmt"

type GoogleAuth struct {
}

func NewGoogleAuth() *GoogleAuth {
	return &GoogleAuth{}
}

func (*GoogleAuth) Login() {
	fmt.Println("google login")
}
