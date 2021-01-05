package eloquent

import (
	"blog/repositories/interfaces"
	"fmt"
)

type PostRepository struct {
	BaseRepo BaseRepository
}

func NewPostRepository() interfaces.PostRepositoryInterface {
	return &PostRepository{}
}

func (*PostRepository) All() {
	fmt.Println("all function in post repository")
}
