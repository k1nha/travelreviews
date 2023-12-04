package util

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type InputToken struct {
	email    string
	username string
}

func CreateToken(i InputToken) (string, error) {
	mapClaims := jwt.MapClaims{
		"username": i.username,
		"email":    i.email,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)

	tokenString, err := token.SignedString(os.Getenv("secret"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
