package domain

import "github.com/golang-jwt/jwt/v4"

type JwtCustomClaims struct {
	ID   string `json:"id"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}