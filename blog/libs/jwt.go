package libs

import (
	"blog/config"
	"blog/helpers"
	"blog/models"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWT struct{}

func NewJWT() *JWT {
	return &JWT{}
}

func (*JWT) CreateToken(user models.User) (string, error) {
	var (
		err error

		jwtExpireTime = helpers.GetValueOrDefault(config.Env["JWT_TTL"], config.Jwt["ttl"])
	)

	//Creating Access Token
	secretKey := helpers.GetValueOrDefault(config.Env["JWT_SECRET"], config.Jwt["secret_key"]) //this should be in an env file

	if secretKey == "" {
		secretKey = "golangapp"
		fmt.Println("JWT empty")
	}

	atClaims := jwt.MapClaims{}

	atClaims["authorized"] = true

	atClaims["user_id"] = user.ID

	atClaims["exp"] = time.Now().Add(time.Minute * time.Duration(jwtExpireTime.(int))).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(secretKey.(string)))

	if err != nil {
		return "", err
	}
	return token, nil
}

func (*JWT) CheckToken() bool {
	return true
}
