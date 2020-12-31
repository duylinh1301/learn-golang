package usercontroller

import (
	"blog/http/response"
	"blog/models"
	"net/http"

	repository "blog/repositories"
)

// AllUser get list users

// Param: http.ResponseWriter, *http.Request

// return Mix

func GetAll(w http.ResponseWriter, r *http.Request) {
	user := []models.User{}

	repository.All(&user)

	response.ReturnJSON(w, http.StatusOK, "", user)
}
