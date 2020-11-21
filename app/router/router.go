package router

import (
	"go-flashcard/app/router/routes"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter()
	return routes.SetupRoutes(r)
}
