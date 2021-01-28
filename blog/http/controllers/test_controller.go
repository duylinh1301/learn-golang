package controllers

import (
	"blog/http/response"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type TestController struct{}

func NewTestController() *TestController {
	return &TestController{}
}

func (testController *TestController) Test(w http.ResponseWriter, r *http.Request) {

	password := "Linh test"

	hashedString, _ := HashPassword(password)

	result := CheckPasswordHash(password+"2", hashedString)

	response.ReturnJSON(w, http.StatusOK, "", map[string]interface{}{
		"result": result,
	})

	return
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
