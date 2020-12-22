package bootstrap

import (
	"log"

	"github.com/joho/godotenv"
)

var env = loadEnv()

//LoadConfig load env and app config
func LoadConfig() map[string]string {
	// var config []string

	envConfig := loadEnv()

	return envConfig
}

func loadEnv() map[string]string {
	var err error

	var data map[string]string

	data, err = godotenv.Read()

	if err != nil {
		log.Println(err)
	}

	return data
}
