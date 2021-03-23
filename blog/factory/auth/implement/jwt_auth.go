package implement

import (
	"blog/helpers"
	"blog/libs"
	"blog/models"
	"blog/repositories/implement"
	"blog/repositories/interfaces"
)

type JWTAuth struct {
	jwt            *libs.JWT
	userRepository interfaces.UserRepositoryInterface
}

func NewJWTAuth() *JWTAuth {
	return &JWTAuth{
		jwt:            libs.NewJWT(),
		userRepository: implement.NewUserRepository(),
	}
}

func (jwtAuth *JWTAuth) Login(data map[string]interface{}) (interface{}, *models.User) {
	email := data["email"]
	password := data["password"]

	user := jwtAuth.userRepository.FirstBy(map[string]interface{}{
		"email": email,
	})

	if *user == (models.User{}) {
		return false, nil
	}

	result := helpers.CheckBrypt(password.(string), user.Password)

	if result == false {
		return false, nil
	}

	accessToken, _ := jwtAuth.jwt.CreateToken(*user)

	return accessToken, user
}
