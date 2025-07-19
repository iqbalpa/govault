package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken() (string, error) {
	key := []byte(os.Getenv("JWT_SECRET_KEY"))
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss":           "govault",
			"exp":           time.Now().Add(time.Minute * 30).Unix(),
			"authenticated": true,
		})
	s, err := t.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("failed to generate token")
	}
	return s, nil
}

func ExtractClaims(tokenStr string) (jwt.MapClaims, error) {
	key := []byte(os.Getenv("JWT_SECRET_KEY"))
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
