// Package routes defines the API routes for the Users-And-Posts-REST service
package routes

import (
	"encoding/json"
	"net/http"
)

// HomeResponse is a struct representing the response data for the home route
type HomeResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

// HomeHandler is an http.HandlerFunc that returns a JSON response with a welcome message
// for the home route
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Set the content type header to JSON
	w.Header().Set("Content-Type", "application/json")

	// Set the HTTP status code to 200 OK
	w.WriteHeader(http.StatusOK)

	// Encode the HomeResponse struct as JSON and write it to the response writer
	json.NewEncoder(w).Encode(HomeResponse{
		Message: "Welcome to Users And Posts REST",
		Status:  true,
	})
}
