package implement

import (
	"blog/helpers"
	"blog/models"
	"blog/repositories/connection"

	"gorm.io/gorm"
)

type UserRepository struct {
	connection *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		connection: connection.NewConnectionDB(),
	}
}

func (userRepository *UserRepository) Register(data *models.User) bool {
	hashedPassword, _ := helpers.HashBcrypt(data.Password)

	data.Password = hashedPassword

	userRepository.connection.Create(data)

	return true
}
