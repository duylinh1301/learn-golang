package implement

import (
	"blog/models"
	"blog/repositories/connection"
	"blog/repositories/interfaces"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	connection *gorm.DB
}

func NewCategoryRepository() interfaces.CategoryRepositoryInterface {
	return &CategoryRepository{
		connection: connection.NewConnectionDB(),
	}
}

func (categoryRepository CategoryRepository) All() *[]models.Category {
	arrayCategory := []models.Category{}

	categoryRepository.connection.Find(&arrayCategory)

	return &arrayCategory
}
