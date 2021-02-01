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

func (userRepository *UserRepository) CreateUserHashPassword(data *models.User) bool {
	hashedPassword, _ := helpers.HashBcrypt(data.Password)

	data.Password = hashedPassword

	userRepository.connection.Create(data)

	return true
}

// FindByID get record by id
func (userRepository *UserRepository) FindByID(id string) *models.User {
	data := models.User{}

	userRepository.connection.First(&data, "id = ?", id)

	return &data
}

// FirstBy get record with condition map[string]interface{}
func (userRepository *UserRepository) FirstBy(condition map[string]interface{}) *models.User {

	user := models.User{}

	userRepository.connection.Select("username, password").Where(condition).First(&user)

	return &user
}
