package postcontroller

import (
	// db "blog/bootstrap"

	"blog/http/response"
	"blog/models"
	"net/http"

	repository "blog/repositories"
)

// GetAll return all posts
func GetAll(w http.ResponseWriter, r *http.Request) {

	post := []models.Post{}

	repository.All(&post)

	response.ReturnJSON(w, http.StatusOK, "", post)
}
