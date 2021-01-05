package models

// Post model mapping posts table
type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
