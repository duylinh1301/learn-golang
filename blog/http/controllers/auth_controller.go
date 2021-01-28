package controllers

import (
	"blog/http/request"
	"blog/http/response"
	"blog/repositories/implement"
	"blog/repositories/interfaces"
	"fmt"
	"net/http"
	"reflect"
)

type AuthController struct {
	userRepository interfaces.UserRepositoryInterface
}

func NewAuthController() *AuthController {
	return &AuthController{
		userRepository: implement.NewUserRepository(),
	}
}

func (authController *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	response.ReturnJSON(w, http.StatusOK, "", map[string]interface{}{
		"access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoX3V1aWQiOiIxZGQ5MDEwYy00MzI4LTRmZjMtYjllNi05NDRkODQ4ZTkzNzUiLCJhdXRob3JpemVkIjp0cnVlLCJ1c2VyX2lkIjo3fQ.Qy8l-9GUFsXQm4jqgswAYTAX9F4cngrl28WJVYNDwtM",
		"user_id":      "1",
	})
}

func (authController *AuthController) Register(w http.ResponseWriter, r *http.Request) {

	// Validate data register

	// Create new user

	data := new(interface{})

	err := request.DecodeJSONBody(r, &data)

	// data

	// authController.userRepository.Register(map[string]interface{}{
	// 	// "username": data.username,
	// })

	s := reflect.ValueOf(*data)
	fmt.Println(reflect.TypeOf(s))
	for _, k := range s.MapKeys() {
		fmt.Println(reflect.TypeOf(k))
		fmt.Println(k, "=>", s.MapIndex(k))
	}

	// val := reflect.TypeOf(*data.email)
	// n := val.FieldByName("Email").Interface().(string)
	// fmt.Println(val)

	// fmt.Println(data["email".(string)])

	fmt.Println(err)

	// Create Jwt
	response.ReturnJSON(w, http.StatusOK, "Register successfully!", map[string]interface{}{
		"access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoX3V1aWQiOiIxZGQ5MDEwYy00MzI4LTRmZjMtYjllNi05NDRkODQ4ZTkzNzUiLCJhdXRob3JpemVkIjp0cnVlLCJ1c2VyX2lkIjo3fQ.Qy8l-9GUFsXQm4jqgswAYTAX9F4cngrl28WJVYNDwtM",
		"user_id":      1,
	})
}
