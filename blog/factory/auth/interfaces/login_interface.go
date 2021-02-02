package interfaces

import "blog/models"

type LoginInterface interface {
	Login(map[string]interface{}) (interface{}, *models.User)
}
