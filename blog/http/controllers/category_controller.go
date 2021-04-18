package controllers

import (
	"blog/http/response"
	"blog/models"
	"blog/repositories/implement"
	"blog/repositories/interfaces"
	"net/http"

	"github.com/gorilla/mux"
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
func (categoryController *CategoryController) Index(w http.ResponseWriter, r *http.Request) {

	categories := categoryController.repository.All()

	response.ReturnJSON(w, http.StatusOK, response.MessageSuccess, categories)

	return
}

func (categoryController *CategoryController) Delete(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	categories := categoryController.repository.FindByID(id)

	if *categories == (models.Category{}) {
		response.ReturnJSON(w, http.StatusOK, response.MessageNotExistsValidate, nil)

		return
	}

	response.ReturnJSON(w, http.StatusOK, response.MessageDeleteSuccess, nil)

	return
}
