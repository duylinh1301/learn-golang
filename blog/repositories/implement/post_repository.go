package implement

import (
	"blog/models"
	"blog/repositories/common"
	"blog/repositories/interfaces"
	"fmt"
)

// PostRepository map post model
type PostRepository struct {
}

var (
	baseRepo   interfaces.BaseRepositoryInterface = common.NewMysqlRepository()
	arrayModel                                    = []models.Post{}
	model                                         = models.Post{}
)

// NewPostRepository create instance
func NewPostRepository() *PostRepository {
	return &PostRepository{}
}

// All get all
func (*PostRepository) All() interface{} {
	baseRepo.All(&arrayModel)
	return &arrayModel
}

func (*PostRepository) Create(data models.BaseModelInterface) {
	// modelIns := models.Post{
	// 	Title:   "Title hardcode",
	// 	Content: "Content hardcode",
	// }

	// table := db.NewScope(model).TableName()
	fmt.Println(model.TableName())

	// baseRepo.Create(map[string]interface{}{
	// 	"Title":   modelIns.Title,
	// 	"Content": modelIns.Content,
	// })
}
