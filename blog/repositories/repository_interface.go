package repositories

type RepositoryInterface interface {
	NewBaseRepository() RepositoryInterface
}
