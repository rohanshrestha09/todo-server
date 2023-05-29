package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/rohanshrestha09/todo/models"
)

func SignJwt(username string) (string, error) {
	expirationTime := time.Now().Add(30 * 1440 * time.Minute)

	claims := &models.Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_TOKEN")))

	return tokenString, err
}

func ParseJwt(jwtToken string) (*models.Claims, *jwt.Token, error) {
	claims := &models.Claims{}

	token, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_TOKEN")), nil
	})

	return claims, token, err
}
