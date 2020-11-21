package routes

import (
	"go-flashcard/app/controllers"
	"net/http"
)

var homeRoutes = []Route{
	Route{
		URI:     "/home",
		Method:  http.MethodGet,
		Handler: controllers.HomeHandler,
	},
}
