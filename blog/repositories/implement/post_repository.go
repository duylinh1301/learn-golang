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

	return &arrayPost
	// return &arrayModelPost
}

// func (*PostRepository) Create(dataPostInterface models.BaseModelInterface) {

// 	dataPost := dataPostInterface.(models.Post)

// 	baseRepo.Create(modelPost.TableName(), map[string]interface{}{
// 		"Title":   dataPost.Title,
// 		"Content": dataPost.Content,
// 	})
// }
