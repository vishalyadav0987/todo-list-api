package authapp

import (
	"time"
)

type CustomClaims struct {
	UserID string
}

type TokenManager interface {
	GenerateToken(userID string, duration time.Duration) (string, error)
	VerifyToken(tokenString string) (*CustomClaims, error)
}
