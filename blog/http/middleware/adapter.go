package middleware

import (
	"net/http"
)

type MiddlewareAdapter func(http.Handler) http.Handler

func Adapt(handler http.Handler, resourceAdapters []MiddlewareAdapter, routeAdapters []MiddlewareAdapter) http.Handler {
	// The loop is reversed so the adapters/middleware gets executed in the same
	// order as provided in the array.
	for i := len(routeAdapters); i > 0; i-- {
		handler = routeAdapters[i-1](handler)
	}

	for i := len(resourceAdapters); i > 0; i-- {
		handler = resourceAdapters[i-1](handler)
	}

	return handler
}
