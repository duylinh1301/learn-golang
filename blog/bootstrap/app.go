package bootstrap

import (
	"blog/config"
	"fmt"
	"log"
	"net/http"

	"gorm.io/gorm"
)

var db *gorm.DB

type Adapter func(http.Handler) http.Handler

// Adapt takes Handler funcs and chains them to the main handler.
func Adapt(handler http.Handler, adapters ...Adapter) http.Handler {
	// The loop is reversed so the adapters/middleware gets executed in the same
	// order as provided in the array.
	for i := len(adapters); i > 0; i-- {
		handler = adapters[i-1](handler)
	}
	return handler
}

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
