package route

import (
	objects "blog/bootstrap/objects"
	postcontroller "blog/controllers/blog"
	"net/http"
)

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
