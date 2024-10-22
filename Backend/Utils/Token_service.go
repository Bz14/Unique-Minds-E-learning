package utils

import (
	"time"
	domain "unique-minds/Domain"
	infrastructures "unique-minds/Infrastructures"

	"github.com/golang-jwt/jwt/v4"
)

type TokenService struct{
	config infrastructures.Config
}

func NewTokenService(config infrastructures.Config) *TokenService {
	return &TokenService{
		config: config,
	}
}


func (ts *TokenService) GenerateAccessToken(user *domain.User)(string, error){
	exp := time.Now().Add(time.Hour * time.Duration(ts.config.Expiry))
	claims := &domain.JwtCustomClaims{
		ID: user.ID.Hex(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	return ts.CreateToken(claims)
}

func (ts *TokenService) GenerateResetToken(user *domain.User)(string, error){
	exp := time.Now().Add(time.Hour * time.Duration(ts.config.Expiry))
	claims := &domain.JwtCustomClaims{
		ID: user.ID.Hex(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	return ts.CreateToken(claims)
}


func (ts *TokenService) CreateToken(claims jwt.Claims)(string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(ts.config.Secret))
	return tokenString, nil
}