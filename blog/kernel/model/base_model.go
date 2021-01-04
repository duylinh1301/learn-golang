package kernel

import (
	"blog/config"
	"fmt"
	"log"
	"reflect"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// BaseModel for CRUD function
type BaseModel struct {
	model interface{}
}

// var model = setModel()
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

func (baseModel *BaseModel) GetInstance() *BaseModel {
	return baseModel
}

func (baseModel *BaseModel) SetModel(model interface{}) {
	fmt.Println("set model function")

	baseModel.model = model
	fmt.Println("Param")
	fmt.Println(reflect.TypeOf(model))
	fmt.Println("Model composition")
	fmt.Println(reflect.TypeOf(baseModel.model))
}

func (baseModel *BaseModel) All() interface{} {
	modelInstance := []baseModel.model

	fmt.Println("all function base model")

	fmt.Println(reflect.TypeOf(modelInstance))

	db.Find(&modelInstance)
	return &modelInstance
}
