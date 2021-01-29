package helpers

import (
	"blog/config"
	"fmt"
	"os"
	"reflect"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken() (string, error) {
	var err error
	//Creating Access Token
	secretKey := os.Getenv("JWT_SECRETs") //this should be in an env file

	fmt.Println(config.Jwt)

	if secretKey == "" {
		secretKey = "golangapp"
		fmt.Println("JWT empty")
	}

	fmt.Println("value jwt secret")
	fmt.Println(reflect.TypeOf(secretKey))

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = 1
	atClaims["exp"] = GetValueOrDefault(config.Env["JWT_TTL"], config.Jwt["ttl"])
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
