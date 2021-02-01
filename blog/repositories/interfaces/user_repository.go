package interfaces

import "blog/models"

type UserRepositoryInterface interface {
	FindByID(string) *models.User
	FirstBy(map[string]interface{}) *models.User
	CreateUserHashPassword(*models.User) bool
}
