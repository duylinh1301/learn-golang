package postcontroller

import (
	response "blog/http/response"
	"blog/models"
	"net/http"
)

// GetAllPosts return all posts
func GetAllPosts(w http.ResponseWriter, r *http.Request) {

	post := models.Post{
		Id:      1,
		Title:   "Lorem ipsum dolor sit amet.",
		Content: "Lonsectetur adipiscing elit. Donec facilisis dapibus suscipit. Duis dignissim eleifend justo, quis consequat lorem lacinia vel.",
	}

	response.ReturnJSON(w, http.StatusOK, "", post)
}
