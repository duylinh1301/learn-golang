package response

import (
	"encoding/json"
	"fmt"
)

type JsonResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (jsonRes *JsonResponse) returnJsonResponse(isSuccess bool, message string) {
	jsonRes = &JsonResponse{
		Success: isSuccess,
		Message: message,
	}

	res, err := json.Marshal(jsonRes)

	if err != nil {
		fmt.Fprintf(response, err.Error())
	}

	fmt.Fprintf(response, string(res))
}
