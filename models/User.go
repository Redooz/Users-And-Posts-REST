package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Id       string `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"unique_index;not null"`
	Password string `json:"password" gorm:"not null"`
	Posts    []Post `json:"posts"`
}
