package dashboard

import (
	"fmt"
	"net/http"
)

func Dashboard(response http.ResponseWriter, request *http.Request) {

	fmt.Fprintf(response, "Dashboard page")
}
