package user

import (
	"blog/http/request"
	"blog/http/response"
	"blog/models"
	"blog/repositories/implement"
	"log"
	"net/http"
)

var (
	userRepository = implement.NewUserRepository()
)

// UserController struct
type UserController struct {
}

// NewUserController contructor
func NewUserController() UserController {
	return UserController{}
}

// GetAll get list users
func (*UserController) GetAll(w http.ResponseWriter, r *http.Request) {
	// user := []models.User{}

	// userRepository.All(&user)

	// response.ReturnJSON(w, http.StatusOK, "", user)

	// user := models.User{}

	// // user.BaseModel.All()

	// user.BaseModel.SetModel(user)
	// user.BaseModel.All()
	return
}

func (*UserController) Create(w http.ResponseWriter, r *http.Request) {

	var post models.Post

	err := request.DecodeJSONBody(r, &post)

	userRepository.Create(post)

	if err != nil {
		log.Println(err.Error())
		response.ReturnJSON(w, http.StatusUnsupportedMediaType, err.Error(), nil)
		return
	}

	response.ReturnJSON(w, http.StatusOK, "", nil)
	return
}
