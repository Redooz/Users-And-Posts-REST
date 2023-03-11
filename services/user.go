package services

import (
	"github.com/Redooz/Users-And-Posts-REST/database"
	"github.com/Redooz/Users-And-Posts-REST/models"
	"gorm.io/gorm"
)

// User is a struct used for interacting with user models in the database.
type User struct {
}

// Create is a function used for creating a new user record in the database.
// It takes a pointer to a User model as an argument and returns a GORM database object.
func (u *User) Create(user *models.User) *gorm.DB {
	var result = database.DB.Create(&user)

	return result
}