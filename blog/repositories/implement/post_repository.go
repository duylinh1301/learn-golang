package implement

import (
	"blog/models"
	"blog/repositories/common"
	"blog/repositories/interfaces"
)

// PostRepository map post model
type PostRepository struct {
}

var (
	baseRepo interfaces.BaseRepositoryInterface = common.NewMysqlRepository()
	model                                       = []models.Post{}
)

// NewPostRepository create instance
func NewPostRepository() *PostRepository {
	return &PostRepository{}
}

// All get all
func (*PostRepository) All() interface{} {
	baseRepo.All(&model)
	return &model
}
