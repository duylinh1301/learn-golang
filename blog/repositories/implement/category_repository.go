package implement

type CategoryRepository struct{}

// NewCategoryRepository constructor
func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{}
}

// All get all
func (*CategoryRepository) All() interface{} {
	baseRepo.All(&arrayModelPost)
	return &arrayModelPost
}
