package bootstrap

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

//LoadConfig load env and app config
func LoadConfig() map[string]string {
	// var config []string

	envConfig := loadEnv()

	return envConfig
}

func loadEnv() map[string]string {
	var err error

	var env map[string]string

	env, err = godotenv.Read()

	fmt.Println(env)

	if err != nil {
		log.Println(err)
	}

	return env
}
