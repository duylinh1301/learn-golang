package interfaces

import "blog/models"

type PostRepositoryInterface interface {
	FindById(string) *models.Post
	All() *[]models.Post
	Create(models.Post)
	Delete(*models.Post)
}
