package postcontroller

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

// GetAll return all posts
func GetAll(w http.ResponseWriter, r *http.Request) {

	post := postRepository.All()

	response.ReturnJSON(w, http.StatusOK, "", post)

	return
}

// Create posts data
func Create(w http.ResponseWriter, r *http.Request) {

	var p models.Post

	err := request.DecodeJSONBody(r, p)

	if err != nil {
		log.Println(err.Error())
		response.ReturnJSON(w, http.StatusUnsupportedMediaType, err.Error(), nil)
		return
	}

	response.ReturnJSON(w, http.StatusOK, "", nil)
	return

	// fmt.Println(r.Header.Get("Content-Type"))

	// decoder := json.NewDecoder(r.Body)
	// var p models.Post
	// err := decoder.Decode(&p)
	// if err != nil {
	// 	panic(err)
	// }
	// log.Println(p)
}
