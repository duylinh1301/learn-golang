package route

import (
	"blog/bootstrap/objects"
	controllers "blog/http/controllers"
	"blog/http/middleware"
	"net/http"
)

// Define array of routes
var (
	// Middleware
	jwtMiddleware        *middleware.JWTMiddleware        = middleware.NewJWTMiddleware()
	permissionMiddleware *middleware.PermissionMiddleware = middleware.NewPermissionMiddleware()

	// Controller
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
					Middleware: []middleware.MiddlewareAdapter{
						jwtMiddleware.CheckJWT(),
					},
				},
			},
		},

		objects.ResourceRoute{
			Prefix: "posts",
			Middleware: []middleware.MiddlewareAdapter{
				jwtMiddleware.CheckJWT(),
			},
			Route: []objects.Route{
				{
					Uri:     "/",
					Name:    "posts.index",
					Method:  http.MethodGet,
					Handler: postcontroller.Index,
					Middleware: []middleware.MiddlewareAdapter{
						permissionMiddleware.CheckRole(),
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
