package postcontroller

import (
	// db "blog/bootstrap"
	response "blog/http/response"
	models "blog/models"
	repository "blog/repositories"
	"net/http"
)

// GetAllPosts return all posts
func GetAllPosts(w http.ResponseWriter, r *http.Request) {

	// post := models.Post{
	// 	Id:      1,
	// 	Title:   "Lorem ipsum dolor sit amet.",
	// 	Content: "Lonsectetur adipiscing elit. Donec facilisis dapibus suscipit. Duis dignissim eleifend justo, quis consequat lorem lacinia vel.",
	// }

	// db := db.ConnectDB()

	post := []models.Post{}

	repository.Query(post)
	// db.Find(&post)

	response.ReturnJSON(w, http.StatusOK, "", post)
}
