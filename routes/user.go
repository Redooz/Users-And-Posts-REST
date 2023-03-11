package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Redooz/Users-And-Posts-REST/models"
	"github.com/Redooz/Users-And-Posts-REST/server"
	"github.com/Redooz/Users-And-Posts-REST/services"
	"github.com/golang-jwt/jwt/v4"
	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
)

var service services.User

// SignUpLoginRequest is the request body for user sign up and login.
type SignUpLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SignUpResponse is the response body for user sign up.
type SignUpResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// SignUpHandler is the HTTP handler function for user sign up.
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var request SignUpLoginRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := ksuid.NewRandom()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := models.User{
		Id:       id.String(),
		Email:    request.Email,
		Password: string(hashedPassword),
	}

	result := service.Create(&user)

	if result.Error != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(SignUpResponse{
		Id:    user.Id,
		Email: user.Email,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// LoginHandler is an http.HandlerFunc that handles user login requests.
// It receives a server instance as parameter and returns an http.HandlerFunc.
func LoginHandler(s server.Server) http.HandlerFunc {
	// Returns an HTTP handler function for handling user login requests
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode login request from JSON payload
		var request SignUpLoginRequest
		err := json.NewDecoder(r.Body).Decode(&request)

		if err != nil {
			// Return error response with status code 400 (Bad Request) if request is invalid
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Get user with the given email from the database
		user, result := service.GetUserByEmail(request.Email)

		if result.Error != nil {
			// Return error response with status code 400 (Bad Request) if user cannot be retrieved from the database
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if user == nil {
			// Return error response with status code 401 (Unauthorized) if user with the given email and password cannot be found
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}

		// Compare the hashed password from the database with the password provided in the login request
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

		if err != nil {
			// Return error response with status code 401 (Unauthorized) if password is incorrect
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}

		// Create JWT token with user ID and expiration time
		claims := models.AppClaim{
			UserId: user.Id,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour * 24)),
			},
		}

		// Sign JWT token with server's secret key
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(s.Config().JWTSecret))

		if err != nil {
			// Return error response with status code 500 (Internal Server Error) if there is an error signing the JWT token
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Return success response with status code 200 (OK) and JWT token in the response body
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(LoginResponse{
			Token: tokenString,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
