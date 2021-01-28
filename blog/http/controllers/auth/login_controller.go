package auth

import (
	"blog/http/response"
	"net/http"
)

type LoginController struct{}

func Login(w http.ResponseWriter, r *http.Request) {
	response.ReturnJSON(w, http.StatusOK, "", map[string]interface{}{
		"token":   "fafdsfsdafdasfdadfd",
		"user_id": "1",
	})
}
