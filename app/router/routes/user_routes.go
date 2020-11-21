package routes

import (
	"go-flashcard/app/controllers"
	"net/http"
)

var userRoutes = []Route{
	Route{
		URI:     "/users/register",
		Method:  http.MethodGet,
		Handler: controllers.UserHandler(controllers.ShowRegister),
	},
	Route{
		URI:     "/users/register",
		Method:  http.MethodPost,
		Handler: controllers.UserHandler(controllers.RegisterHandler),
	},
	Route{
		URI:     "/users/login",
		Method:  http.MethodGet,
		Handler: controllers.UserHandler(controllers.ShowLogin),
	},
	Route{
		URI:     "/users/login",
		Method:  http.MethodPost,
		Handler: controllers.UserHandler(controllers.LoginHandler),
	},
}
