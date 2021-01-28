package route

import (
	"blog/bootstrap/objects"
	controllers "blog/http/controllers"
	"net/http"
)

// Define array of routes
var (
	// testcontroller     *controllers.TestController    = controllers.NewTestController()
	authcontroller     *controllers.AuthController    = controllers.NewAuthController()
	postcontroller     controllers.PostController     = controllers.NewPostController()
	categorycontroller controllers.CategoryController = controllers.NewCategoryController()
	// usercontroller     user.UserController     = user.NewUserController()

	// ApiRoutes = objects.NewGroupRoute()

	ApiRoutes = []objects.GroupRoute{

		objects.GroupRoute{
			Prefix: "auth",
			Route: []objects.Route{
				{
					Name:    "auth.register",
					Uri:     "/register",
					Method:  http.MethodPost,
					Handler: authcontroller.Register,
				},
				{
					Name:    "auth.login",
					Uri:     "/login",
					Method:  http.MethodPost,
					Handler: authcontroller.Login,
				},
			},
		},

		objects.GroupRoute{
			Prefix: "posts",
			Route: []objects.Route{
				{
					Uri:     "/",
					Name:    "posts.index",
					Method:  http.MethodGet,
					Handler: postcontroller.GetAll,
				},
				{
					Uri:     "/",
					Name:    "posts.create",
					Method:  http.MethodPost,
					Handler: postcontroller.Create,
				},
				{
					Uri:     "/{id}",
					Method:  http.MethodPut,
					Handler: postcontroller.Update,
				},
				{
					Uri:     "/{id}",
					Method:  http.MethodDelete,
					Handler: postcontroller.Delete,
				},
			},
		},

		objects.GroupRoute{
			Prefix: "categories",
			Route: []objects.Route{
				{
					Name:    "categories.index",
					Uri:     "/",
					Method:  http.MethodGet,
					Handler: categorycontroller.GetAll,
				},
			},
		},
	}
)
