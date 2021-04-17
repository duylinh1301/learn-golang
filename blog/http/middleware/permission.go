package middleware

import (
	"fmt"
	"net/http"
)

type PermissionMiddleware struct{}

func NewPermissionMiddleware() *PermissionMiddleware {
	return &PermissionMiddleware{}
}

func (permissionMiddleware PermissionMiddleware) CheckRole() MiddlewareAdapter {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("CheckRole function")

			// Serve the next handler
			next.ServeHTTP(w, r)
		})
	}
}
