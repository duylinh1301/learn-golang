package models

import "net/http"

type Post struct {
	id      int    `json:"id"`
	title   string `json:"title"`
	content string `json:"content"`
}

// Get all posts
func getAllPosts(w http.ResponseWriter, r *http.Request) {

}
