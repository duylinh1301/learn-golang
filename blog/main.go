package main

import (
	"blog/route"
	"log"
	"net/http"
)

func handleRequests() {
	// Declare router
	route.Router()

	// Start the server on port 10000
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
