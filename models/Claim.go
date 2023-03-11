package models

import "github.com/golang-jwt/jwt/v4"

type AppClaim struct {
	UserId string `json:"userId"`
	jwt.RegisteredClaims
}
