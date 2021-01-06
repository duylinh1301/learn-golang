package eloquent

import (
	"blog/repositories"
	"blog/repositories/interfaces"
	"fmt"
)

type PostRepository struct {
	repositories.BaseRepository
}

func NewPostRepository() interfaces.PostRepositoryInterface {
	return &PostRepository{}
}

func (*PostRepository) All() {
	fmt.Println("all function in post eloquent repository")
}
