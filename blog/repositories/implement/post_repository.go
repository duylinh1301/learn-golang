package implement

import (
	"blog/models"
)

// PostRepository map post model
type PostRepository struct {
}

var (
	arrayModelPost = []models.Post{}
	modelPost      = models.Post{}
)

// NewPostRepository create instance
func NewPostRepository() *PostRepository {

	return &PostRepository{}
}

// All get all
func (*PostRepository) All() interface{} {
	baseRepo.All(&arrayModelPost)
	return &arrayModelPost
}

func (*PostRepository) Create(dataPostInterface models.BaseModelInterface) {

	dataPost := dataPostInterface.(models.Post)

	baseRepo.Create(modelPost.TableName(), map[string]interface{}{
		"Title":   dataPost.Title,
		"Content": dataPost.Content,
	})
}
