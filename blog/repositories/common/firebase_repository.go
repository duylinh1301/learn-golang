package common

import (
	"blog/repositories/interfaces"
	"fmt"
)

// FirebaseRepository connect to MYSQL database
type FirebaseRepository struct {
}

// var db = connectDB()
// var env = config.Env

// func connectDB() *gorm.DB {
// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", env["DB_USERNAME"], env["DB_PASSWORD"], env["DB_HOST"], env["DB_PORT"], env["DB_DATABASE"])

// 	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Println("Connection Failed to Open")
// 	} else {
// 		log.Println("Connection to " + env["DB_CONNECTION"] + " established")
// 	}

// 	return connection
// }

// NewFirebaseRepository create instance
func NewFirebaseRepository() interfaces.BaseRepositoryInterface {
	return &FirebaseRepository{}
}

// All function get all records from model
func (*FirebaseRepository) All(model interface{}) {
	fmt.Println("All function in Firebase repo")
}
