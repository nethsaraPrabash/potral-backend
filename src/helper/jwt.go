package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID uint, role string, secretKey string) (string, error) {
	claims := jwt.MapClaims{
		"exp":  jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
		"iat":  jwt.NewNumericDate(time.Now()),
		"sub":  userID,
		"role": role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
