package implement

import (
	"blog/models"
	"fmt"
)

type UserRepository struct{}

var (
	modelUser      = models.User{}
	arrayModelUser = []models.User{}
)

// NewPostRepository create instance
func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (*UserRepository) All() {
	baseRepo.All(&modelUser)
}

func (*UserRepository) Create(data models.BaseModelInterface) {

	// modelIns := models.Post{
	// 	Title:   "Title hardcode",
	// 	Content: "Content hardcode",
	// }

	// table := db.NewScope(model).TableName()
	fmt.Println(modelUser.TableName())

	// baseRepo.Create(map[string]interface{}{
	// 	"Title":   modelIns.Title,
	// 	"Content": modelIns.Content,
	// })
}
