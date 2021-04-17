package bootstrap

import (
	"blog/bootstrap/objects"
	"blog/http/middleware"
	"blog/route"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// InitRoute init route
func InitRoute() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	return setupRoutes(r)
}

// LoadResouceRoute load all route defined in route/api.go
func LoadResouceRoute() []objects.ResourceRoute {
	routes := route.ApiRoutes

	return routes
}

func setupRoutes(r *mux.Router) *mux.Router {
	for _, resourceRoute := range LoadResouceRoute() {
		prefix := resourceRoute.Prefix

		for _, route := range resourceRoute.Route {
			numberResourceMiddleware := len(resourceRoute.Middleware)

			numberRouteMiddleware := len(route.Middleware)

			if numberResourceMiddleware > 0 || numberRouteMiddleware > 0 {
				r.HandleFunc("/api/"+prefix+trimRouteUrl(route.Uri), middleware.Adapt(
					http.HandlerFunc(route.Handler),
					resourceRoute.Middleware,
					route.Middleware,
				).ServeHTTP).Methods(route.Method)

			} else {
				r.HandleFunc("/api/"+prefix+trimRouteUrl(route.Uri), route.Handler).Methods(route.Method)
			}

		}

	}

	return r
}

func trimRouteUrl(url string) string {
	result := strings.TrimRight(url, "/")

	return result
}
