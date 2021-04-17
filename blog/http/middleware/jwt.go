package middleware

import (
	"blog/helpers"
	"blog/http/response"
	"blog/repositories/implement"
	jwtsupport "blog/support/jwtauth"
	"net/http"
)

type JWTMiddleware struct {
	userRepository *implement.UserRepository
}

func NewJWTMiddleware() *JWTMiddleware {
	return &JWTMiddleware{
		userRepository: implement.NewUserRepository(),
	}
}

func (jwtMiddleware JWTMiddleware) CheckJWT() MiddlewareAdapter {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			bearerToken := r.Header.Get("Authorization")

			jwtSupportInstant := jwtsupport.NewJWT()

			err := jwtSupportInstant.VerifyToken(bearerToken)

			if err != nil {
				response.ReturnJSON(w, http.StatusUnauthorized, err.Error(), nil)

				return
			}

			jwtClaim := jwtSupportInstant.ExtractClaims(bearerToken)

			user := jwtMiddleware.userRepository.FindByID(jwtClaim["user_id"])

			helpers.SetUser(user)

			// Serve the next handler
			next.ServeHTTP(w, r)
		})
	}
}
