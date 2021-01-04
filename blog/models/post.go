package models

import kernel "blog/kernel/model"

// Post model
type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`

	// BaseModel kernel.BaseModel

	BaseModel kernel.BaseModel
}
