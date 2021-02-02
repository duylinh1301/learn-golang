package bootstrap

import (
	"blog/bootstrap/objects"
	"blog/route"
	"fmt"
	"strings"

	"github.com/gorilla/mux"
)

// InitRoute init route
func InitRoute() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return setupRoutes(r)
}

// LoadRoutes load all route defined in route/api.go
func LoadRouteGroups() []objects.GroupRoute {
	routes := route.ApiRoutes

	return routes
}
func setupRoutes(r *mux.Router) *mux.Router {
	for _, routeGroups := range LoadRouteGroups() {
		prefix := routeGroups.Prefix

		for _, route := range routeGroups.Route {
			fmt.Println(prefix, route.Uri, route.Method)
			r.HandleFunc("/api/"+prefix+trimRouteUrl(route.Uri), route.Handler).Methods(route.Method)
		}

	}

	return r
}

func trimRouteUrl(url string) string {
	result := strings.TrimRight(url, "/")

	return result
}
