package route

import (
	objects "blog/bootstrap/objects"
	controllers "blog/http/controllers"
	"net/http"
)

// Define array of routes
var (
	postcontroller     controllers.PostController     = controllers.NewPostController()
	categorycontroller controllers.CategoryController = controllers.NewCategoryController()
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
			Method:  http.MethodPut,
			Handler: postcontroller.Update,
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
