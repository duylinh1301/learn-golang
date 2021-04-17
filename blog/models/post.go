package models

import (
	"time"
)

// Post model mapping posts table
type Post struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	Category_id int       `json:"category_id"`
	Category    Category  `gorm:"foreignKey:Category_id"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}

func NewPost() *Post {
	return &Post{
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}
}

func (Post) TableName() string {
	return "posts"
}

func (post *Post) GetField() []map[string]interface{} {
	// fields := map[string]interface{
	// 	{"ID" : post.ID},
	// 	{"Title" : post.Title},
	// 	{"Content" : post.Content},
	// }

	return []map[string]interface{}{
		{"Name": "jinzhu_1", "Age": 18},
		{"Name": "jinzhu_2", "Age": 20},
	}

	// return {
	// 	"ID" : post.ID,
	// 	"Title" : post.Title,
	// 	"Content" : post.Content,
	// }
}
