package routes

import (
	"go-flashcard/app/config"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

func Load() []Route {
	routes := make([]Route, 0)
	routes = append(routes, homeRoutes...)
	routes = append(routes, userRoutes...)
	return routes
}

func LoadStatic(r *mux.Router) {
	fs := http.FileServer(http.Dir("./app" + config.STATIC_DIR))
	r.PathPrefix(config.STATIC_DIR).Handler(http.StripPrefix(config.STATIC_DIR, fs))
}

func SetupRoutes(r *mux.Router) *mux.Router {
	//Load static
	LoadStatic(r)
	//Load routes
	for _, route := range Load() {
		r.HandleFunc(route.URI, route.Handler).Methods(route.Method)
	}
	return r
}
