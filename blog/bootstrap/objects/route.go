package objects

import (
	"blog/middleware"
	"net/http"
)

type Route struct {
	Name              string
	Uri               string
	Method            string
	MiddlewareHandler []middleware.MiddlewareAdapter
	Handler           func(http.ResponseWriter, *http.Request)
}
