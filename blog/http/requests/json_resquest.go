package requests

import (
	exceptions "blog/exceptions/requests"
	"encoding/json"
	"net/http"
)

const contentType string = "application/json"

// DecodeJSONBody decode http request body to model data
func DecodeJSONBody(request *http.Request, model interface{}) error {
	if request.Header.Get("Content-Type") != contentType {
		msg := "Content-Type header is not application/json"

		return &exceptions.MalformedException{Status: http.StatusUnsupportedMediaType, Msg: msg}
	}

	decoder := json.NewDecoder(request.Body)

	decoder.DisallowUnknownFields()

	err := decoder.Decode(&model)

	if err != nil {
		msg := err.Error()

		return &exceptions.MalformedException{Status: http.StatusBadRequest, Msg: msg}
	}

	return nil
}
