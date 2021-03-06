package route

import (
	"blog/bootstrap/objects"
	controllers "blog/http/controllers"
	"blog/middleware"
	"net/http"
)

// Define array of routes
var (
	// testcontroller     *controllers.TestController    = controllers.NewTestController()
	authcontroller     *controllers.AuthController    = controllers.NewAuthController()
	postcontroller     controllers.PostController     = controllers.NewPostController()
	categorycontroller controllers.CategoryController = controllers.NewCategoryController()
	// usercontroller     user.UserController     = user.NewUserController()

	// ApiRoutes = objects.NewResourceRoute()

	ApiRoutes = []objects.ResourceRoute{

		objects.ResourceRoute{
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
				{
					Name:    "auth.logout",
					Uri:     "/logout",
					Method:  http.MethodPost,
					Handler: authcontroller.Logout,
				},
			},
		},

		objects.ResourceRoute{
			Prefix: "posts",
			Middleware: []middleware.MiddlewareAdapter{
				middleware.CheckJWT(),
			},
			Route: []objects.Route{
				{
					Uri:     "/",
					Name:    "posts.index",
					Method:  http.MethodGet,
					Handler: postcontroller.GetAll,
					Middleware: []middleware.MiddlewareAdapter{
						middleware.CheckRole(),
					},
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

		objects.ResourceRoute{
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
