package implement

import (
	"blog/repositories/common"
	"blog/repositories/interfaces"
	"fmt"
)

// PostRepository map post model
type PostRepository struct {
}

var (
	baseRepo interfaces.BaseRepositoryInterface = common.NewMysqlRepository()
)

// NewPostRepository create instance
func NewPostRepository() *PostRepository {
	return &PostRepository{}
}

// All get all
func (*PostRepository) All() {
	fmt.Println("Ham All cua post repository")
	baseRepo.All()
}
