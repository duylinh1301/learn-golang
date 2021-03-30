package middleware

import (
	"blog/http/response"
	jwtsupport "blog/support/jwtauth"
	"fmt"
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

func CheckJWT() MiddlewareAdapter {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			bearerToken := r.Header.Get("Authorization")

			err := jwtsupport.NewJWT().VerifyToken(bearerToken)

			if err != nil {
				response.ReturnJSON(w, http.StatusUnauthorized, err.Error(), nil)

				return
			}

			// Serve the next handler
			next.ServeHTTP(w, r)
		})
	}
}

func CheckRole() MiddlewareAdapter {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("CheckRole function")

			// Serve the next handler
			next.ServeHTTP(w, r)
		})
	}
}
