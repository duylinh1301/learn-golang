package bootstrap

import (
	"blog/bootstrap/objects"
	"blog/route"

	"github.com/gorilla/mux"
)

// InitRoute init route
func InitRoute() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	return SetupRoutes(r)
}

// LoadRoutes load all route defined in route/api.go
func LoadRoutes() []objects.Route {
	routes := route.ApiRoutes

	return routes
}

// SetupRoutes set route for listen server
func SetupRoutes(r *mux.Router) *mux.Router {
	for _, route := range LoadRoutes() {
		r.HandleFunc("/api"+route.Uri, route.Handler).Methods(route.Method)
	}

	return r
}
