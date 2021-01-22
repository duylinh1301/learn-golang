package blog

import (
	"fmt"
	"net/http"
)

type CategoryController struct{}

func NewCategoryController() CategoryController {
	return CategoryController{}
}

func (*CategoryController) GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all function")
}
