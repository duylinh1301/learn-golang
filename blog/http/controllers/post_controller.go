package controllers

import (
	// db "blog/bootstrap"

	"blog/http/requests"
	postrequest "blog/http/requests/post"
	"blog/http/response"
	"blog/models"
	"blog/repositories/implement"
	"blog/repositories/interfaces"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
func (postController *PostController) Index(w http.ResponseWriter, r *http.Request) {

	post := postController.postRepository.All()

	response.ReturnJSON(w, http.StatusOK, "", post)

	return
}

// Create posts data
func (postController *PostController) Create(w http.ResponseWriter, r *http.Request) {

	//

	postRequest := postrequest.NewCreateRequest()

	err := requests.DecodeJSONBody(r, &postRequest)

	if err != nil {
		log.Println(err.Error())

		response.ReturnJSON(w, http.StatusUnsupportedMediaType, err.Error(), nil)

		return
	}

	// Validate data

	var post = postController.postRepository.GetModel()

	post.Title = postRequest.Title
	post.Content = postRequest.Content
	post.Description = postRequest.Description

	postController.postRepository.Create(*post)

	response.ReturnJSON(w, http.StatusCreated, "", nil)

	return
}

// Update update post data
func (postController *PostController) Update(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	post := postController.postRepository.FindByID(id)

	if (models.Post{}) == *post {
		response.ReturnJSON(w, http.StatusOK, "Post not found!", nil)

		return
	}

	updateRequest := postrequest.NewUpdateRequest()

	err := requests.DecodeJSONBody(r, updateRequest)

	fmt.Println(*updateRequest)

	if err != nil {
		log.Println(err.Error())

		response.ReturnJSON(w, http.StatusUnsupportedMediaType, err.Error(), nil)

		return
	}

	post.Title = updateRequest.Title
	post.Description = updateRequest.Description
	post.Content = updateRequest.Content

	postController.postRepository.UpdateById(id, *post)

	if (models.Post{}) == *post {
		response.ReturnJSON(w, http.StatusNotFound, "Post not found!", nil)

		return
	}

	response.ReturnJSON(w, http.StatusOK, "", nil)

	return
}

// Delete function
func (postController *PostController) Delete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id := vars["id"]

	post := postController.postRepository.FindByID(id)

	if (models.Post{}) == *post {
		response.ReturnJSON(w, http.StatusNotFound, "Post not found!", nil)

		return
	}

	postController.postRepository.Delete(post)

	response.ReturnJSON(w, http.StatusOK, "", nil)

	return
}
