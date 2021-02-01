package libs

import (
	"blog/config"
	"blog/helpers"
	"blog/models"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func JWTCreateToken(user models.User) (string, error) {
	var err error
	//Creating Access Token
	secretKey := helpers.GetValueOrDefault(config.Env["JWT_SECRET"], config.Jwt["secret_key"]) //this should be in an env file

	fmt.Println(config.Jwt)

	if secretKey == "" {
		secretKey = "golangapp"
		fmt.Println("JWT empty")
	}

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = user.ID
	atClaims["exp"] = helpers.GetValueOrDefault(config.Env["JWT_TTL"], config.Jwt["ttl"])
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(secretKey.(string)))
	if err != nil {
		return "", err
	}
	return token, nil
}
