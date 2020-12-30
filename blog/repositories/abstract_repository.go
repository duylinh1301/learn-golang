package repositories

import (
	"fmt"
	"log"

	config "blog/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// AbstractRepository eloquent repository to reuse function query db
type AbstractRepository struct {
}

var db = connectDB()
var env = config.Env

func connectDB() *gorm.DB {
	fmt.Println("2")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", env["DB_USERNAME"], env["DB_PASSWORD"], env["DB_HOST"], env["DB_PORT"], env["DB_DATABASE"])

	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection to " + env["DB_CONNECTION"] + " established")
	}

	return connection
}

// All function get all records from model
func All(model *[]interface{}) interface{} {
	db.Find(&model)
	return model
}
