package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route is the standard struct of routes API
type Route struct {
	URI            string
	Method         string
	Func           func(http.ResponseWriter, *http.Request)
	AuthIsRequired bool
}

// Config puts all the routes in the router
func Config(r *mux.Router) *mux.Router {
	routes := userRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Func).Methods(route.Method)
	}

	return r
}
