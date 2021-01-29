package controllers

import (
	"blog/config"
	"blog/helpers"
	"blog/http/request"
	"blog/http/response"
	"blog/models"
	"blog/repositories/implement"
	"blog/repositories/interfaces"
	"fmt"
	"log"
	"net/http"
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
	// token, _ := CreateToken()

	// fmt.Println(token)

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

	if err != nil {
		log.Println(err.Error())

		response.ReturnJSON(w, http.StatusUnsupportedMediaType, err.Error(), nil)

		return
	}

	mapData := helpers.InterfaceToMap(*data)

	user := models.User{
		Username: mapData["username"].(string),
		Email:    mapData["email"].(string),
		Password: mapData["password"].(string),
	}

	result := authController.userRepository.Register(&user)

	if result == false {
		response.ReturnJSON(w, http.StatusInternalServerError, "Cannot create user!", nil)

		return
	}

	timeToLive := helpers.GetValueOrDefault(config.Env["JWT_TTL"], config.Jwt["ttl"])

	fmt.Println(timeToLive)

	// Create Jwt

	// Return Success message and JWT
	response.ReturnJSON(w, http.StatusOK, "Register successfully!", map[string]interface{}{
		"access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoX3V1aWQiOiIxZGQ5MDEwYy00MzI4LTRmZjMtYjllNi05NDRkODQ4ZTkzNzUiLCJhdXRob3JpemVkIjp0cnVlLCJ1c2VyX2lkIjo3fQ.Qy8l-9GUFsXQm4jqgswAYTAX9F4cngrl28WJVYNDwtM",
		"user_id":      1,
	})
}
