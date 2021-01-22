package blog

import (
	// db "blog/bootstrap"

	"blog/http/request"
	"blog/http/response"
	"blog/models"
	"blog/repositories/implement"
	"log"
	"net/http"
)

var (
	postRepository = implement.NewPostRepository()
	// postRepository interfaces.PostRepositoryInterface = mongodb.NewPostRepository()
)

// PostController struct
type PostController struct{}

// NewPostController contructor
func NewPostController() PostController {
	return PostController{}
}

// GetAll return all posts
func (*PostController) GetAll(w http.ResponseWriter, r *http.Request) {

	post := postRepository.All()

	response.ReturnJSON(w, http.StatusOK, "", post)

	return
}

// Create posts data
func (*PostController) Create(w http.ResponseWriter, r *http.Request) {

	var post models.Post

	err := request.DecodeJSONBody(r, &post)

	postRepository.Create(post)

	if err != nil {
		log.Println(err.Error())
		response.ReturnJSON(w, http.StatusUnsupportedMediaType, err.Error(), nil)
		return
	}

	response.ReturnJSON(w, http.StatusOK, "", nil)
	return
}
