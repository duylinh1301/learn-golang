package models

type Category struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewCategory() *Category {
	return &Category{}
}
