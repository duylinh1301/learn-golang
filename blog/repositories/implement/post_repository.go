package implement

import (
	"blog/models"
	"blog/repositories/connection"
	"blog/repositories/interfaces"
	"fmt"
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
	// return &arrayModelPost
}

// Create post function
func (postRepository *PostRepository) Create(data models.Post) {

	createdEnitity := postRepository.connection.Create(&data)

	fmt.Println(createdEnitity)

	return
}

// Update post function
func (postRepository *PostRepository) UpdateById(id string, data models.Post) {
	data.Updated_at = time.Now()
	postRepository.connection.Where("id = ?", id).Updates(&data)
}

// Delete post function
func (postRepository *PostRepository) Delete(data *models.Post) {

	postRepository.connection.Delete(data)

	return
}
