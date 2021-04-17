package implement

import (
	"blog/models"
	"blog/repositories/connection"
	"blog/repositories/interfaces"
	"time"

	"gorm.io/gorm"
)

// PostRepository map post model
type PostRepository struct {
	connection *gorm.DB
}

// NewPostRepository create instance
func NewPostRepository() interfaces.PostRepositoryInterface {
	return &PostRepository{
		connection: connection.NewConnectionDB(),
	}
}

func (postRepository PostRepository) GetModel() *models.Post {
	return models.NewPost()
}

func (postRepository PostRepository) FindByID(id string) *models.Post {
	data := models.Post{}

	postRepository.connection.First(&data, "id = ?", id)

	return &data
}

// All get all
func (postRepository PostRepository) All() *[]models.Post {
	arrayPost := []models.Post{}

	postRepository.connection.Find(&arrayPost)

	return &arrayPost
}

// Create post function
func (postRepository *PostRepository) Create(data models.Post) {
	postRepository.connection.Create(&data)

	return
}

// Update post function
func (postRepository *PostRepository) UpdateById(id string, data models.Post) {
	data.Updated_at = time.Now()

	postRepository.connection.Where("id = ?", id).Updates(&data)

	return
}

// Delete post function
func (postRepository *PostRepository) Delete(data *models.Post) {
	postRepository.connection.Delete(data)

	return
}
