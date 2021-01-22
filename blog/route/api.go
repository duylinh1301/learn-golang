package route

import (
	objects "blog/bootstrap/objects"
	blog "blog/http/controllers/blog"
	usercontroller "blog/http/controllers/user"
	"net/http"
)

// Define array of routes
var (
	postcontroller     blog.PostController     = blog.NewPostController()
	categorycontroller blog.CategoryController = blog.NewCategoryController()

	ApiRoutes = []objects.Route{
		objects.Route{
			Uri:     "/posts",
			Method:  http.MethodGet,
			Handler: postcontroller.GetAll,
		},
		objects.Route{
			Uri:     "/posts",
			Method:  http.MethodPost,
			Handler: postcontroller.Create,
		},
		objects.Route{
			Uri:     "/categories",
			Method:  http.MethodGet,
			Handler: categorycontroller.GetAll,
		},
		objects.Route{
			Uri:     "/users",
			Method:  http.MethodGet,
			Handler: usercontroller.GetAll,
		},
	}
)
