package middleware

import (
	"blog/http/response"
	jwtsupport "blog/support/jwtauth"
	"net/http"
)

type JWTMiddleware struct {
}

func NewJWTMiddleware() *JWTMiddleware {
	return &JWTMiddleware{}
}

func (jwtMiddleware JWTMiddleware) CheckJWT() MiddlewareAdapter {
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
