package objects

import "blog/middleware"

type ResourceRoute struct {
	Prefix     string
	Route      []Route
	Middleware []middleware.MiddlewareAdapter
}
