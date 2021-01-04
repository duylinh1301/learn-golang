package bootstrap

import (
	"fmt"
	"log"
	"net/http"

	config "blog/config"

	"gorm.io/gorm"
)

var db *gorm.DB

// HandleRequests load all config for every request
func HandleRequests() {

	// Load application config
	env := config.Env

	// Declare router
	route := InitRoute()

	// Start the server on port
	port := "10000"

	value, exists := env["API_PORT"]

	if exists {
		port = value
	}

	domain, exists := env["DB_HOST"]

	fmt.Printf("[Listening]: %s:%s\n", domain, port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), route))
}
