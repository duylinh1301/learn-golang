package request

import (
	exceptions "blog/exceptions/requests"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const contentType string = "application/json"

func DecodeJSONBody(request *http.Request, model interface{}) error {
	if request.Header.Get("Content-Type") != contentType {
		msg := "Content-Type header is not application/json"

		return &exceptions.MalformedException{Status: http.StatusUnsupportedMediaType, Msg: msg}
	}

	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()

	fmt.Println(decoder)

	err := decoder.Decode(&model)
	if err != nil {
		panic(err)
	}
	log.Println(model)

	// decoder := json.NewDecoder(r.Body)
	// var p models.Post
	// err := decoder.Decode(&model)
	// if err != nil {
	// 	panic(err)
	// }
	// log.Println(p)

	return nil
}
