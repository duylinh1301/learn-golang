package controllers

import (
	"blog/http/response"
	"blog/repositories/implement"
	"blog/repositories/interfaces"
	"net/http"
)

// CategoryController struct
type CategoryController struct {
	repository interfaces.CategoryRepositoryInterface
}

// NewCategoryController contructor
func NewCategoryController() CategoryController {
	return CategoryController{
		repository: implement.NewCategoryRepository(),
	}
}

// GetAll return all posts
func (categoryController *CategoryController) GetAll(w http.ResponseWriter, r *http.Request) {

	categories := categoryController.repository.All()

	response.ReturnJSON(w, http.StatusOK, "", categories)

	return

}
