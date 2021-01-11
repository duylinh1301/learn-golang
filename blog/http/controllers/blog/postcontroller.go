package postcontroller

import (
	// db "blog/bootstrap"

	"blog/http/response"
	"blog/repositories/implement"
	"net/http"
)

var (
	postRepository = implement.NewPostRepository()
	// postRepository interfaces.PostRepositoryInterface = mongodb.NewPostRepository()
)

// GetAll return all posts
func GetAll(w http.ResponseWriter, r *http.Request) {

	post := postRepository.All()

	response.ReturnJSON(w, http.StatusOK, "", post)
}
