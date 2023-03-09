// Package routes defines the API routes for the Users-And-Posts-REST service
package routes

import (
	"github.com/Redooz/Users-And-Posts-REST/server"
	"github.com/gorilla/mux"
)

// BindRouter sets up the HTTP routing for the Users-And-Posts-REST service.
// It takes a Server interface and a mux.Router as arguments and maps the appropriate
// HTTP handlers to their respective routes using the router.
func BindRouter(s server.Server, r *mux.Router) {
	r.HandleFunc("/api/v1", HomeHandler(s))
}
