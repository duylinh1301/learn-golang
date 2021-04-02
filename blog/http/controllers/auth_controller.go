package controllers

import (
	"blog/factory/auth"
	"blog/helpers"
	"blog/http/request"
	userrequest "blog/http/request/user"
	"blog/http/response"
	"blog/repositories/implement"
	"blog/repositories/interfaces"
	jwtsupport "blog/support/jwtauth"
	"log"
	"net/http"
)

type AuthController struct {
	maxAttempts    int
	jwt            *jwtsupport.JWT
	userRepository interfaces.UserRepositoryInterface
}

func NewAuthController() *AuthController {
	return &AuthController{
		maxAttempts:    3,
		jwt:            jwtsupport.NewJWT(),
		userRepository: implement.NewUserRepository(),
	}
}

func (authController *AuthController) Login(w http.ResponseWriter, r *http.Request) {

	// Get data Body
	data := new(interface{})

	request.DecodeJSONBody(r, &data)

	mapData := helpers.InterfaceToMap(*data)

	condition := map[string]interface{}{
		"email":    mapData["email"],
		"password": mapData["password"],
	}

	authFactory := auth.NewAuthFactory(auth.TypeJWT)

	accessToken, user := authFactory.AuthMethod.Login(condition)

	if accessToken == false {
		response.ReturnJSON(w, http.StatusUnauthorized, "Username or password incorrect!", nil)

		return
	}

	response.ReturnJSON(w, http.StatusOK, "", map[string]interface{}{
		"access_token": accessToken,
		"user": map[string]interface{}{
			"username": user.Username,
			"email":    user.Email,
		},
	})

	return
}

func (authController *AuthController) Register(w http.ResponseWriter, r *http.Request) {

	// Validate data register
	userRequest := userrequest.NewUserRequest()

	err := request.DecodeJSONBody(r, &userRequest)

	if err != nil {
		log.Println(err.Error())

		response.ReturnJSON(w, http.StatusUnsupportedMediaType, err.Error(), nil)

		return
	}

	// Create new user

	user := authController.userRepository.GetModel()

	user.Username = userRequest.Username
	user.Email = userRequest.Email
	user.Password = userRequest.Password

	result := authController.userRepository.CreateUserHashPassword(user)

	if result == false {
		response.ReturnJSON(w, http.StatusInternalServerError, "Cannot create user!", nil)

		return
	}

	// Return Success message and JWT
	response.ReturnJSON(w, http.StatusOK, "Register successfully!", map[string]interface{}{
		"user": map[string]interface{}{
			"username": user.Username,
			"email":    user.Email,
		},
	})

	return
}

func (authController *AuthController) Logout(w http.ResponseWriter, r *http.Request) {

	reqToken := r.Header.Get("Authorization")

	authController.jwt.AddToBlackList(reqToken)

	response.ReturnJSON(w, http.StatusResetContent, "Logout successfully!", nil)
}
