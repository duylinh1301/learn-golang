package interfaces

import "blog/models"

type CategoryRepositoryInterface interface {
	All() *[]models.Category
}
