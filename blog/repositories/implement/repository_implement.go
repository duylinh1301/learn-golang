package implement

import (
	"blog/repositories"
)

type RepositoryImplement struct {
}

func (baseRepo *RepositoryImplement) NewBaseRepository() repositories.RepositoryInterface {
	return &RepositoryImplement{}
}
