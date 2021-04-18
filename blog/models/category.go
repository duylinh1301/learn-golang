package models

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Posts       []Post
}

func NewCategory() *Category {
	return &Category{}
}
