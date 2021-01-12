package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSONResponse define return data format
type JSONResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

// ReturnJSON return json data type for every request
func ReturnJSON(w http.ResponseWriter, code int, message string, data interface{}) {

	if message == "" {
		message = "Success!"
	}

	jsonRes := JSONResponse{
		StatusCode: code,
		Message:    message,
		Data:       data,
	}

	jsonData, err := json.Marshal(jsonRes)

	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(jsonRes.StatusCode)

	w.Write(jsonData)

	return
}
