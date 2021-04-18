package models

import (
	"time"
)

// Post model mapping posts table
type Post struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	CategoryID  int    `json:"category_id"`
	User_id     int    `json:"user_id"`
	Category    Category
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}

func NewPost() *Post {
	return &Post{
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}
}
