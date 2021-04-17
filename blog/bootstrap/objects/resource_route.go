package objects

import "blog/http/middleware"

type ResourceRoute struct {
	Prefix     string
	Route      []Route
	Middleware []middleware.MiddlewareAdapter
}
