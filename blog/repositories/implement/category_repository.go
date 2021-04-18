package implement

import (
	"blog/models"
	"blog/repositories/connection"
	"blog/repositories/interfaces"
	"fmt"

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

func (categoryRepository CategoryRepository) FindByID(id interface{}) *models.Category {
	arrayCategory := models.Category{}

	categoryRepository.connection.Where("id", "=", id).Find(&arrayCategory)

	return &arrayCategory
}

func (categoryRepository CategoryRepository) All() *[]models.Category {
	arrayCategory := []models.Category{}

	categoryRepository.connection.Find(&arrayCategory)

	return &arrayCategory
}

func (categoryRepository CategoryRepository) Delete(data *models.Category) {
	categoryRepository.connection.Delete(data)

	return
}

func (categoryRepository CategoryRepository) CountPostsByID(id interface{}) int64 {

	var totalPosts int64 = 0

	fmt.Println(totalPosts)

	// PostReposit

	return totalPosts
}
