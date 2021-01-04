package postcontroller

import (
	// db "blog/bootstrap"

	"blog/http/response"
	"blog/models"
	"net/http"
)

// GetAll return all posts
func GetAll(w http.ResponseWriter, r *http.Request) {

	// postRepo := repositories.PostsRepository{}

	// postRepo.BaseRepository.All()

	post := models.Post{}

	// post.BaseModel.All()

	post.BaseModel.SetModel(post)
	data := post.BaseModel.All()

	// fmt.Println(reflect.TypeOf(post))

	// repository.All(&post)
	// postRepo := repository.PostsRepository{}

	// postRepo.BaseRepository.All(&post)

	response.ReturnJSON(w, http.StatusOK, "", data)
}
