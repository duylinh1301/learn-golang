package interfaces

import "blog/models"

type UserRepositoryInterface interface {
	Register(*models.User)
}
