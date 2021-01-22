package blog

import (
	"fmt"
	"net/http"
)

// CategoryController struct
type CategoryController struct{}

// NewCategoryController contructor
func NewCategoryController() CategoryController {
	return CategoryController{}
}

// GetAll return all posts
func (*CategoryController) GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all function")
}
