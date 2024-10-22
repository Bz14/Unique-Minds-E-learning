package domain

import "github.com/golang-jwt/jwt/v4"

type JwtCustomClaims struct {
	ID   string `json:"id"`
	jwt.RegisteredClaims
}

type TokenServiceInterface interface{
	CreateToken(claims jwt.Claims)(string, error)
	GenerateAccessToken(user *User)(string, error)
	GenerateResetToken(user *User)(string, error)
}