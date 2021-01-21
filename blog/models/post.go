package models

// Post model mapping posts table
type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
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
