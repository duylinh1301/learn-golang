package interfaces

type BaseRepositoryInterface interface {
	All(model interface{})
	Create(table string, data map[string]interface{})
}
