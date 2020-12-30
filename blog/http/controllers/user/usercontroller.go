package usercontroller

import (
	"blog/http/response"
	"blog/models"
	"net/http"
)

// AllUser get list users
func GetAll(w http.ResponseWriter, r *http.Request) {
	user := models.User{
		Id:       1,
		Username: "Linh Nguyen",
		Email:    "linhnguyen@okxe.vn",
	}

	response.ReturnJSON(w, http.StatusOK, "", user)
}
