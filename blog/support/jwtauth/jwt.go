package jwtauth

import (
	"blog/config"
	"blog/helpers"
	"blog/models"
	supportcache "blog/support/cache"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
	storage string
}

func NewJWT() *JWT {
	return &JWT{
		storage: "storage/blacklist.txt",
	}
}

func (jwtStruct *JWT) CreateToken(user models.User) (string, error) {
	var (
		err error
	)

	//Creating Access Token
	secretKey := helpers.GetValueOrDefault(config.Env["JWT_SECRET"], config.Jwt.SecretKey) //this should be in an env file

	if secretKey == "" {
		secretKey = "golangapp"
		fmt.Println("JWT empty")
	}

	atClaims := jwt.MapClaims{}

	atClaims["authorized"] = true

	atClaims["user_id"] = user.ID

	atClaims["exp"] = time.Now().Add(time.Minute * time.Duration(config.Jwt.Ttl)).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(secretKey.(string)))

	if err != nil {
		return "", err
	}
	return token, nil
}

func (jwtStruct *JWT) VerifyToken(tokenString string) error {

	secretKey := helpers.GetValueOrDefault(config.Env["JWT_SECRET"], config.Jwt.SecretKey) //this should be in an env file

	tokenString = jwtStruct.ExtractToken(tokenString)

	cache := supportcache.NewCacheFile()

	_, found := cache.Get(tokenString)

	if found {
		return fmt.Errorf("this token has been logout!")
	}

	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey.(string)), nil
	})

	return err
}

func (jwtStruct *JWT) ExtractToken(bearerToken string) string {
	splitedToken := strings.Split(bearerToken, "Bearer ")

	if len(splitedToken) == 2 {
		return splitedToken[1]
	}

	return ""
}

func (jwtStruct *JWT) AddToBlackList(tokenString string) {

	tokenString = jwtStruct.ExtractToken(tokenString)

	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		log.Fatal(err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Fatalf("Can't convert token's claims to standard claims")
	}

	var tm time.Time
	switch exp := claims["exp"].(type) {
	case float64:
		tm = time.Unix(int64(exp), 0)
	case json.Number:
		v, _ := exp.Int64()
		tm = time.Unix(v, 0)
	}

	leftTime := time.Now().Sub(tm)

	cache := supportcache.NewCacheFile()

	cache.Set(tokenString, time.Now().Unix(), leftTime)
}

func (jwtStruct *JWT) ExtractClaims(tokenString string) jwt.MapClaims {

	tokenString = jwtStruct.ExtractToken(tokenString)

	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		fmt.Println(1)
		log.Fatal(err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Fatalf("Can't convert token's claims to standard claims")
	}

	return claims
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
