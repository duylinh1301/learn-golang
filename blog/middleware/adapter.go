package middleware

import (
	"blog/http/response"
	"blog/libs"
	"fmt"
	"net/http"
)

type MiddlewareAdapter func(http.Handler) http.Handler

func Adapt(handler http.Handler, adapters []MiddlewareAdapter) http.Handler {
	// The loop is reversed so the adapters/middleware gets executed in the same
	// order as provided in the array.
	for i := len(adapters); i > 0; i-- {
		handler = adapters[i-1](handler)
	}
	return handler
}

func CheckJWT() MiddlewareAdapter {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("CheckJWT function")

			bearerToken := r.Header.Get("Authorization")

			err := libs.NewJWT().VerifyToken(bearerToken)

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
