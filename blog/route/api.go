package route

import (
	objects "blog/bootstrap/objects"
	blog "blog/http/controllers/blog"
	"net/http"
)

// Define array of routes
var (
	postcontroller     blog.PostController     = blog.NewPostController()
	categorycontroller blog.CategoryController = blog.NewCategoryController()
	// usercontroller     user.UserController     = user.NewUserController()

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
			Uri:     "/posts/{id}",
			Method:  http.MethodDelete,
			Handler: postcontroller.Delete,
		},
		// objects.Route{
		// 	Uri:     "/users",
		// 	Method:  http.MethodGet,
		// 	Handler: usercontroller.GetAll,
		// },
		// objects.Route{
		// 	Uri:     "/users",
		// 	Method:  http.MethodPost,
		// 	Handler: usercontroller.Create,
		// },
	}
)
