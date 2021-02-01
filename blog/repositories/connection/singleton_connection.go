package connection

import (
	"blog/config"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var connection *gorm.DB

func NewConnectionDB() *gorm.DB {

	if connection == nil {

		var env = config.Env

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", env["DB_USERNAME"], env["DB_PASSWORD"], env["DB_HOST"], env["DB_PORT"], env["DB_DATABASE"])

		gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Println("Connection Failed to Open")
		} else {
			log.Println("Connection to " + env["DB_CONNECTION"] + " established")
		}

		connection = gormDB

	}

	return connection
}
