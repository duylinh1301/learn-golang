package postcontroller

import (
	// db "blog/bootstrap"

	"blog/repositories/interfaces"
	"blog/repositories/mongodb"
	"fmt"
	"net/http"
	"reflect"
)

var (
	// postRepository interfaces.PostRepositoryInterface = eloquent.NewPostRepository()
	postRepository interfaces.PostRepositoryInterface = mongodb.NewPostRepository()
)

// GetAll return all posts
func GetAll(w http.ResponseWriter, r *http.Request) {

	fmt.Println(reflect.TypeOf(postRepository))

	// postRepo := eloquent.NewPostRepository()

	fmt.Println(postRepository)

	// postRepository.BaseRepo.Test()

	postRepository.BaseRepo.Test()

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
