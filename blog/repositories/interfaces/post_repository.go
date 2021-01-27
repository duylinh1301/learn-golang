package interfaces

import "blog/models"

type PostRepositoryInterface interface {
	All() *[]models.Post
	Create(models.Post)
}
