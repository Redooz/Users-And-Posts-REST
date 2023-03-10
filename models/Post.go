package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model

	Id      string `json:"id" gorm:"primaryKey"`
	Content string `json:"content"`
	UserId  uint   `json:"userId"`
}
