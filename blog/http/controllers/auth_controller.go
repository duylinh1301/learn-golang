package controllers

import (
	"blog/http/request"
	"blog/http/response"
	"blog/libs"
	"blog/models"
	"blog/repositories/implement"
	"blog/repositories/interfaces"
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
	response.ReturnJSON(w, http.StatusOK, "", map[string]interface{}{
		"access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoX3V1aWQiOiIxZGQ5MDEwYy00MzI4LTRmZjMtYjllNi05NDRkODQ4ZTkzNzUiLCJhdXRob3JpemVkIjp0cnVlLCJ1c2VyX2lkIjo3fQ.Qy8l-9GUFsXQm4jqgswAYTAX9F4cngrl28WJVYNDwtM",
		"user_id":      "1",
	})
}

func (authController *AuthController) Register(w http.ResponseWriter, r *http.Request) {

	// Validate data register

	data := new(interface{})

	err := request.DecodeJSONBody(r, &data)

	if err != nil {
		log.Println(err.Error())

		response.ReturnJSON(w, http.StatusUnsupportedMediaType, err.Error(), nil)

		return
	}

	// Create new user

	// mapData := helpers.InterfaceToMap(*data)

	// inputUser := models.User{
	// 	Username: mapData["username"].(string),
	// 	Email:    mapData["email"].(string),
	// 	Password: mapData["password"].(string),
	// }

	// result := authController.userRepository.Register(&inputUser)

	// if result == false {
	// 	response.ReturnJSON(w, http.StatusInternalServerError, "Cannot create user!", nil)

	// 	return
	// }

	createdUser := models.User{
		ID:       15,
		Username: "linhnguyen",
		Email:    "linhnguyen@okxe.vn",
	}

	// Create Jwt
	accessToken, err := libs.JWTCreateToken(createdUser)

	if err != nil {
		response.ReturnJSON(w, http.StatusInternalServerError, "Cannot create token!", nil)

		return
	}

	// Return Success message and JWT
	response.ReturnJSON(w, http.StatusOK, "Register successfully!", map[string]interface{}{
		"access_token": accessToken,
		"user": map[string]interface{}{
			"id":       createdUser.ID,
			"username": createdUser.Username,
			"email":    createdUser.Email,
		},
	})

	return
}
