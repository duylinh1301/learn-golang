package interfaces

import "blog/models"

type PostRepositoryInterface interface {
	GetModel() *models.Post
	FindByID(string) *models.Post
	All() *[]models.Post
	Create(models.Post)
	UpdateById(id string, data models.Post)
	Delete(*models.Post)
}
