package services

import (
	"github.com/Redooz/Users-And-Posts-REST/database"
	"github.com/Redooz/Users-And-Posts-REST/models"
	"gorm.io/gorm"
)

type User struct {
}

func (u *User) Create(user *models.User) *gorm.DB{
	var result = database.DB.Create(&user)

	return result
}
