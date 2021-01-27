package blog

import (
	// db "blog/bootstrap"

	"blog/http/request"
	"blog/http/response"
	"blog/models"
	"blog/repositories/implement"
	"blog/repositories/interfaces"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	postRepository = implement.NewPostRepository()
	// postRepository interfaces.PostRepositoryInterface = mongodb.NewPostRepository()
)

// PostController struct
type PostController struct {
	postRepository interfaces.PostRepositoryInterface
}

// NewPostController contructor
func NewPostController() PostController {
	return PostController{
		postRepository: implement.NewPostRepository(),
	}
}

// GetAll return all posts
func (postController *PostController) GetAll(w http.ResponseWriter, r *http.Request) {

	post := postController.postRepository.All()

	response.ReturnJSON(w, http.StatusOK, "", post)

	return
}

// Create posts data
func (postController *PostController) Create(w http.ResponseWriter, r *http.Request) {

	var post models.Post

	err := request.DecodeJSONBody(r, &post)

	// Validate data

	if err != nil {
		log.Println(err.Error())

		response.ReturnJSON(w, http.StatusUnsupportedMediaType, err.Error(), nil)

		return
	}

	postController.postRepository.Create(post)

	response.ReturnJSON(w, http.StatusOK, "", nil)

	return
}

// Delete function
func (postController *PostController) Delete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id := vars["id"]

	post := postController.postRepository.FindById(id)

	if (models.Post{}) == *post {
		response.ReturnJSON(w, http.StatusNotFound, "Post not found!", nil)

		return
	}

	postController.postRepository.Delete(post)

	response.ReturnJSON(w, http.StatusOK, "", nil)

	return
}
