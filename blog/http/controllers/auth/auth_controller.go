package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// Login use for login
func Login(response http.ResponseWriter, request *http.Request) {
	jsonRes := JsonResponse{
		Success: true,
		Message: "Login successful!",
	}

	fmt.Print(jsonRes)

	res, err := json.Marshal(jsonRes)

	if err != nil {
		fmt.Fprintf(response, err.Error())
	}

	fmt.Fprintf(response, string(res))
}

// Logout use for logout
func Logout(response http.ResponseWriter, request *http.Request) {
	jsonRes := JsonResponse{
		Success: true,
		Message: "Logout successful!",
	}

	fmt.Print(jsonRes)

	res, err := json.Marshal(jsonRes)

	if err != nil {
		fmt.Fprintf(response, err.Error())
	}

	fmt.Fprintf(response, string(res))
}
