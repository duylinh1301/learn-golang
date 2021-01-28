package common

import (
	"blog/config"
	"blog/repositories/interfaces"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MysqlRepository connect to MYSQL database
type MysqlRepository struct {
}

var db = connectDB()
var env = config.Env

func connectDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", env["DB_USERNAME"], env["DB_PASSWORD"], env["DB_HOST"], env["DB_PORT"], env["DB_DATABASE"])

	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection to " + env["DB_CONNECTION"] + " established")
	}

	return connection
}

// NewMysqlRepository create instance
func NewMysqlRepository() interfaces.BaseRepositoryInterface {
	return &MysqlRepository{}
}

// All function get all records from model
func (*MysqlRepository) All(model interface{}) {
	db.Find(model)
}

// Create function create a record from model
func (*MysqlRepository) Create(table string, data map[string]interface{}) {
	db.Table("posts").Create(data)
}
