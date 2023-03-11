package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Redooz/Users-And-Posts-REST/models"
	"github.com/Redooz/Users-And-Posts-REST/services"
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
	json.NewEncoder(w).Encode(SignUpResponse{
		Id:    user.Id,
		Email: user.Email,
	})

}
