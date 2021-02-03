package objects

import "blog/middleware"

type ResourceRoute struct {
	Prefix            string
	Route             []Route
	MiddlewareHandler []middleware.MiddlewareAdapter
}
