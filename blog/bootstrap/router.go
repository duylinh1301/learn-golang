package bootstrap

import (
	"blog/bootstrap/objects"
	"blog/middleware"
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
			numberResourceMiddleware := len(resourceRoute.MiddlewareHandler)

			numberRouteMiddleware := len(route.MiddlewareHandler)

			if numberResourceMiddleware > 0 || numberRouteMiddleware > 0 {
				if numberResourceMiddleware > 0 {
					r.HandleFunc("/api/"+prefix+trimRouteUrl(route.Uri), middleware.Adapt(
						http.HandlerFunc(route.Handler),
						resourceRoute.MiddlewareHandler,
					).ServeHTTP).Methods(route.Method)
				}

				if numberRouteMiddleware > 0 {
					r.HandleFunc("/api/"+prefix+trimRouteUrl(route.Uri), middleware.Adapt(
						http.HandlerFunc(route.Handler),
						route.MiddlewareHandler,
					).ServeHTTP).Methods(route.Method)

				}

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
