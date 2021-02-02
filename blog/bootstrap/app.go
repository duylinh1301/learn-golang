package bootstrap

import (
	"blog/config"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var db *gorm.DB

// HandleRequests load all config for every request
func HandleRequests() {

	// Load application config
	env := config.Env

	// Declare router
	// route := InitRoute()

	route := mux.NewRouter().StrictSlash(true)

	route.HandleFunc("/api/test/1", test1).Methods(http.MethodGet)
	route.HandleFunc("/api/test/2", test2).Methods(http.MethodGet)
	route.Use(middlewareTest)

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

func test1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("route test 1")
}

func test2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("route test 2")
}

func middlewareTest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(r)

		// Do stuff here
		log.Println(r.RequestURI)
		fmt.Println("test middleware")
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
