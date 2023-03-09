// Package server provides a simple HTTP server implementation that uses the Gorilla Mux router.
package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Config defines the configuration options for the server.
type Config struct {
	Port        string // Port number to listen on
	JWTSecret   string // Secret key for JWT token validation
	DatabaseURL string // URL of the database connection
}

// Server defines an interface for accessing the server's configuration.
type Server interface {
	Config() *Config // Returns the server configuration
}

// Broker is the main server implementation. It contains the server configuration and router, and provides methods for starting and stopping the server.
type Broker struct {
	config *Config     // Server configuration
	router *mux.Router // HTTP router
}

// Config returns the server configuration.
func (b *Broker) Config() *Config {
	return b.config
}

// Starts the server and listens for incoming requests. It takes a callback function that 
// defines the routes and middleware for the API using the router and server configuration.
func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	// Create a new Gorilla Mux router
	b.router = mux.NewRouter()

	// Bind the routes and middleware to the router
	binder(b, b.router)

	// Log a message indicating that the server is starting
	log.Println("Starting server on port", b.config.Port)

	// Start the HTTP server and listen for incoming requests
	err := http.ListenAndServe(b.config.Port, b.router)

	// If an error occurs, log the error and exit the application
	if err != nil {
		log.Println("error starting server:", err)
	} else {
		log.Fatalf("server stopped")
	}
}

// NewServer creates a new instance of the Broker with the specified configuration. It returns an error if the configuration is invalid.
func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	// Validate the configuration
	err := validateConfig(config)
	if err != nil {
		return nil, err
	}

	// Create a new instance of the Broker with the validated configuration
	broker := Broker{
		config: config,
		router: mux.NewRouter(),
	}

	return &broker, nil
}

// validateConfig checks if the configuration is valid by ensuring that all required fields are present.
func validateConfig(config *Config) error {
	if config.Port == "" {
		return errors.New("port is required")
	}
	if config.JWTSecret == "" {
		return errors.New("jwt secret is required")
	}
	if config.DatabaseURL == "" {
		return errors.New("database url is required")
	}
	return nil
}
