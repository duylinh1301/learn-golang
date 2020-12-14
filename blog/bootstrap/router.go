package bootstrap

import (
	"blog/bootstrap/objects"
	"blog/route"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return SetupRoutes(r)
}

func Load() []objects.Route {
	routes := route.ApiRoutes
	return routes
}

func SetupRoutes(r *mux.Router) *mux.Router {

	for _, route := range Load() {
		r.HandleFunc("/api"+route.Uri, route.Handler).Methods(route.Method)
	}

	return r
}
