package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	authapp "github.com/vishalyadav0987/todo-list-api/internal/application/auth"
)

type jwtClaims struct {
	UserId string `json:"sub"`
	jwt.RegisteredClaims
}

type JWTManager struct {
	secretKey string
}

func NewJWTManager(secret string) *JWTManager {
	return &JWTManager{
		secretKey: secret,
	}
}

func (j *JWTManager) GenerateToken(
	userId string,
	duration time.Duration,
) (string, error) {
	now := time.Now()

	claims := jwtClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(duration)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(j.secretKey))
}

func (j *JWTManager) VerifyToken(tokenString string) (*authapp.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&jwtClaims{},
		func(t *jwt.Token) (any, error) {
			return []byte(j.secretKey), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwtClaims)

	if !ok || !token.Valid {
		return nil, err
	}

	return &authapp.CustomClaims{
		UserID: claims.UserId,
	}, nil
}
