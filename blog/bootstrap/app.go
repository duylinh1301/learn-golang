package bootstrap

import (
	"log"
	"net/http"
)

func HandleRequests() {
	// Declare router
	New()
	// Start the server on port 10000
	log.Fatal(http.ListenAndServe(":10000", nil))
}