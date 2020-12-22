package repositories

import (
	connect "blog/bootstrap"

	"gorm.io/gorm"
)

var db *gorm.DB

func Query(models interface{}) interface{} {
	db = connect.ConnectDB()
	db.Find(&models)
	return models
}
