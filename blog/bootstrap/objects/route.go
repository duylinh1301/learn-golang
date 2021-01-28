package objects

import "net/http"

type Route struct {
	Name    string
	Uri     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

func NewRoute() *Route {
	return &Route{}
}
