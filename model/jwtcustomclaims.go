package model

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type JWTCustomClaims struct {
	UserID uuid.UUID `json:"user_id"`
	Email  string    `json:"email"`
	jwt.StandardClaims
}
