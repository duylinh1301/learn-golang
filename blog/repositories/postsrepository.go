package repositories

import "blog/models"

// PostsRepository post repo
type PostsRepository struct {
	BaseRepository BaseRepository
	Model          models.BaseModel
}

func (postRepo *PostsRepository) GetModel() models.BaseModel {
	return postRepo.Model
}
