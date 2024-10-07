package models

//Data decodifcada en los tokens

import "github.com/golang-jwt/jwt"

type AppClaims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}
