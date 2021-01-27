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
