package config

import (
	"log"

	"github.com/joho/godotenv"
)

var Env = LoadEnv()

func LoadEnv() map[string]string {
	var err error

	var data map[string]string

	data, err = godotenv.Read()

	if err != nil {
		log.Println(err)
	}

	return data
}
