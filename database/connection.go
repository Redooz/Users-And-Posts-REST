package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection(DSN string) {
	var err error
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("DB connected")

}
