module github.com/Redooz/Users-And-Posts-REST

go 1.19

require github.com/gorilla/mux v1.8.0

require github.com/joho/godotenv v1.5.1

require (
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	gorm.io/driver/mysql v1.4.7
	gorm.io/gorm v1.24.6
)

require (
	github.com/golang-jwt/jwt/v4 v4.5.0
	github.com/segmentio/ksuid v1.0.4
	golang.org/x/crypto v0.7.0
)
