package route

import (
	objects "blog/bootstrap/objects"
	postcontroller "blog/controllers/blog"
	"net/http"
)

// Define array of routes
var ApiRoutes = []objects.Route{
	objects.Route{
		Uri:     "/posts",
		Method:  http.MethodGet,
		Handler: postcontroller.GetAllPosts,
	},
	// bootstrap.Route{
	// 	Uri:     "posts",
	// 	Method:  http.MethodGet,
	// 	Handler: postcontroller.GetAllPosts,
	// },
}
