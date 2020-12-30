package route

import (
	objects "blog/bootstrap/objects"
	postcontroller "blog/http/controllers/blog"
	usercontroller "blog/http/controllers/user"
	"net/http"
)

// Define array of routes
var ApiRoutes = []objects.Route{
	objects.Route{
		Uri:     "/posts",
		Method:  http.MethodGet,
		Handler: postcontroller.GetAll,
	},
	objects.Route{
		Uri:     "/users",
		Method:  http.MethodGet,
		Handler: usercontroller.GetAll,
	},
}
