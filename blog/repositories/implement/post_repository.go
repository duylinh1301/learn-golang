package implement

import (
	"blog/models"
	"blog/repositories/connection"
	"blog/repositories/interfaces"

	"gorm.io/gorm"
)

// PostRepository map post model
type PostRepository struct {
	connection *gorm.DB
	model      *models.Post
}

// NewPostRepository create instance
func NewPostRepository() interfaces.PostRepositoryInterface {
	connection := connection.ConnectDB()
	model := new(models.Post)
	return &PostRepository{
		connection,
		model,
	}
}

func (postRepository PostRepository) FindById(id string) *models.Post {

	data := models.Post{}

	postRepository.connection.First(&data, "id = ?", id)

	return &data
}

// All get all
func (postRepository PostRepository) All() *[]models.Post {

	arrayPost := []models.Post{}

	postRepository.connection.Find(&arrayPost)

	return &arrayPost
	// return &arrayModelPost
}

// Create post function
func (postRepository *PostRepository) Create(data models.Post) {

	postRepository.connection.Create(&data)

	return
}

// Delete post function
func (postRepository *PostRepository) Delete(data *models.Post) {

	postRepository.connection.Delete(data)

	return
}
