package interfaces

import "blog/models"

type UserRepositoryInterface interface {
	GetModel() *models.User
	FindByID(interface{}) *models.User
	FirstBy(map[string]interface{}) *models.User
	CreateUserHashPassword(*models.User) bool
	IsEmailExists(string) bool
}
