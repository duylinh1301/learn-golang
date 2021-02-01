package controllers

import (
	"blog/factory/auth"
	"blog/helpers"
	"blog/http/request"
	"blog/http/response"
	"blog/libs"
	"blog/models"
	"blog/repositories/implement"
	"blog/repositories/interfaces"
	"fmt"
	"log"
	"net/http"
)

type AuthController struct {
	maxAttempts    int
	jwt            *libs.JWT
	userRepository interfaces.UserRepositoryInterface
}

func NewAuthController() *AuthController {
	return &AuthController{
		maxAttempts:    3,
		jwt:            libs.NewJWT(),
		userRepository: implement.NewUserRepository(),
	}
}

func (authController *AuthController) Login(w http.ResponseWriter, r *http.Request) {

	// Get data Body
	data := new(interface{})

	request.DecodeJSONBody(r, &data)

	mapData := helpers.InterfaceToMap(*data)

	condition := map[string]interface{}{
		"username": mapData["username"],
	}

	authFactory := auth.NewAuthFactory()

	authFactory.AuthMethod.Login()

	fmt.Println(condition)

	// Find user auth
	user := authController.userRepository.FirstBy(condition)

	fmt.Println(user)

	// Create Jwt
	// accessToken, err := authController.jwt.CreateToken(*user)

	// if err != nil {
	// 	response.ReturnJSON(w, http.StatusInternalServerError, "Cannot create token!", nil)

	// 	return
	// }

	// response.ReturnJSON(w, http.StatusOK, "", map[string]interface{}{
	// 	"access_token": accessToken,
	// 	"user": map[string]interface{}{
	// 		"username": user.Username,
	// 		"email":    user.Email,
	// 	},
	// })

	fmt.Println("done")

	return
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
	mapData := helpers.InterfaceToMap(*data)

	User := models.User{
		Username: mapData["username"].(string),
		Email:    mapData["email"].(string),
		Password: mapData["password"].(string),
	}

	result := authController.userRepository.CreateUserHashPassword(&User)

	if result == false {
		response.ReturnJSON(w, http.StatusInternalServerError, "Cannot create user!", nil)

		return
	}

	// Return Success message and JWT
	response.ReturnJSON(w, http.StatusOK, "Register successfully!", map[string]interface{}{
		"user": map[string]interface{}{
			"username": User.Username,
			"email":    User.Email,
		},
	})

	return
}
