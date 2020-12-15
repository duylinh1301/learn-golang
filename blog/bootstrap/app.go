package bootstrap

import (
	"fmt"
	"log"
	"net/http"
)

// HandleRequests load all config for every request
func HandleRequests() {
	// Declare router
	route := InitRoute()

	// Load application config
	config := LoadConfig()

	// Start the server on port
	port := "10000"

	value, exists := config["API_PORT"]

	if exists {
		port = value
	}

	fmt.Printf("[Listening]: %s\n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), route))
}
