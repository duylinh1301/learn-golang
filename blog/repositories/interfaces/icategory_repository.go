package interfaces

import "blog/models"

type CategoryRepositoryInterface interface {
	All() *[]models.Category
	FindByID(interface{}) *models.Category
	Delete(data *models.Category)
}
