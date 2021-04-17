package controllers

import (
	"blog/factory/auth"
	"blog/helpers"
	"blog/http/requests"
	authrequest "blog/http/requests/auth"
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
	loginRequest := authrequest.NewLoginRequest()

	requests.DecodeJSONBody(r, &loginRequest)

	condition := map[string]interface{}{
		"email":    loginRequest.Email,
		"password": loginRequest.Password,
	}

	authFactory := auth.NewAuthFactory(auth.TypeJWT)

	accessToken, user := authFactory.AuthMethod.Login(condition)

	if accessToken == false {
		response.ReturnJSON(w, http.StatusUnauthorized, response.MessageIncorrectCredentailValidate, nil)

		return
	}

	response.ReturnJSON(w, http.StatusOK, response.MessageSuccess, map[string]interface{}{
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
	registerRequest := authrequest.NewRegisterRequest()

	err := requests.DecodeJSONBody(r, &registerRequest)

	if err != nil {
		log.Println(err.Error())

		response.ReturnJSON(w, http.StatusUnsupportedMediaType, err.Error(), nil)

		return
	}

	// Validate unique email
	if authController.userRepository.IsEmailExists(registerRequest.Email) {

		response.ReturnJSON(w, http.StatusUnprocessableEntity, response.MessageEmailExistsValidate, nil)

		return
	}

	// Create new user

	user := authController.userRepository.GetModel()

	user.Username = registerRequest.Username
	user.Email = registerRequest.Email
	user.Password = registerRequest.Password

	result := authController.userRepository.CreateUserHashPassword(user)

	if result == false {
		response.ReturnJSON(w, http.StatusInternalServerError, "Cannot create user!", nil)

		return
	}

	// Return Success message and JWT
	response.ReturnJSON(w, http.StatusOK, response.MessageRegisterSuccees, map[string]interface{}{
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

	response.ReturnJSON(w, http.StatusResetContent, response.MessageLoginSuccees, nil)

	return
}

func (authController *AuthController) User(w http.ResponseWriter, r *http.Request) {

	response.ReturnJSON(w, http.StatusResetContent, response.MessageSuccess, helpers.AuthUser())

	return
}
