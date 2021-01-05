package mongodb

import (
	"blog/repositories/interfaces"
	"fmt"
)

type PostRepository struct {
}

func NewPostRepository() interfaces.PostRepositoryInterface {
	return &PostRepository{}
}

func (*PostRepository) All() {
	fmt.Println("all function in post repository")
}
