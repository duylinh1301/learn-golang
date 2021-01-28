package implement

import "blog/models"

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (userController *UserRepository) Register(*models.User) {

}
