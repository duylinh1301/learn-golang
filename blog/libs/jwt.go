package libs

import (
	"blog/config"
	"blog/helpers"
	"blog/models"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/patrickmn/go-cache"
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

func (jwtStruct *JWT) VerifyToken(tokenString string) error {

	secretKey := helpers.GetValueOrDefault(config.Env["JWT_SECRET"], config.Jwt["secret_key"]) //this should be in an env file

	tokenString = jwtStruct.extractToken(tokenString)

	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey.(string)), nil
	})

	return err
}

func extractClaims(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecretString := "72XaXFub8lXXOR5MnjqaRQxJXZ4U5569267kP6lGqwRsif7hMqYlX4KCh1Q3jXau" // Value
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}

func (jwtStruct *JWT) AddToBlackList(tokenString string) {

	tokenString = jwtStruct.extractToken(tokenString)

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

	fmt.Println(reflect.TypeOf(tm), tm)

	left := time.Now().Sub(tm)

	c := cache.New(left, left)

	c.Set(tokenString, time.Now().Unix(), cache.DefaultExpiration)

	foo, found := c.Get(tokenString)

	fmt.Println(foo, found)

	// str := fmt.Sprintf("%v", claim["exp"].(string))

	// fmt.Println(str)

	// strconv.ParseFloat()

	// claims := jwt.MapClaims{}

	// secretKey := helpers.GetValueOrDefault(config.Env["JWT_SECRET"], config.Jwt["secret_key"])

	// str := fmt.Sprintf("%v", secretKey)
	// fmt.Println(str) // "[1 2 3]"

	// // fmt.Println((string)secretKey)

	// token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
	// 	return str, nil
	// })
	// // ... error handling

	// expireTimeToken := 0

	// // do something with decoded claims
	// for key, value := range claims {
	// 	if key == "exp" {
	// 		expireTimeToken = fmt.Sprintf("%d", value)
	// 	}
	// }

	// fmt.Println(expireTimeToken)
	// fmt.Println(token, err)
	// c := cache.New(5*time.Minute, 10*time.Minute)

	// c.Set("foo", "bar", cache.DefaultExpiration)

	// foo, found := c.Get("foo")

	// fmt.Println(foo, found)

	// if _, err := os.Stat(jwtStruct.storage); err == nil {

	// 	return err
	// }

	// file, err := os.Create(jwtStruct.storage)

	// ache, err := ristretto.NewCache(&ristretto.Config{
	// 	NumCounters: 1e7,     // Num keys to track frequency of (10M).
	// 	MaxCost:     1 << 30, // Maximum cost of cache (1GB).
	// 	BufferItems: 64,      // Number of keys per Get buffer.
	// })

	// check(err)

	// data := []byte(tokenString + "\n")

	// _, err = file.Write(data)

	// defer file.Close()

	// check(err)

	// return found
}

func (jwtStruct *JWT) extractToken(bearerToken string) string {
	splitedToken := strings.Split(bearerToken, "Bearer ")

	if len(splitedToken) == 2 {
		return splitedToken[1]
	}

	return ""
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
