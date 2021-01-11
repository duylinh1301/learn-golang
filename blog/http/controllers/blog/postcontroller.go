package postcontroller

import (
	// db "blog/bootstrap"

	"blog/repositories/implement"
	"fmt"
	"net/http"
)

var (
	postRepository = implement.NewPostRepository()
	// postRepository interfaces.PostRepositoryInterface = mongodb.NewPostRepository()
)

// GetAll return all posts
func GetAll(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Day la ham GetAll PostController")

	postRepository.All()

	// postRepository.All()

	// post := models.Post{}

	// post.BaseModel.All()

	// post.BaseModel.SetModel(post)
	// data := post.BaseModel.All()

	// fmt.Println(reflect.TypeOf(post))

	// repository.All(&post)
	// postRepo := repository.PostsRepository{}

	// postRepo.BaseRepository.All(&post)

	// response.ReturnJSON(w, http.StatusOK, "", data)
}
