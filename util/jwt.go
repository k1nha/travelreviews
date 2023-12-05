package util

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type InputToken struct {
	Email    string
	Username string
}

func CreateToken(i InputToken) (string, error) {
	mapClaims := jwt.MapClaims{
		"username": i.Username,
		"email":    i.Email,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// func ValidateToken(tkn string) error {
// 	token, err := jwt.Parse(tkn, func(t *jwt.Token) (interface{}, error) {
// 		_, ok := t.Method.(*jwt.SigningMethodECDSA)

// 		if !ok {
// 			writer.WriteHeader(http.StatusUnauthorized)
// 			_, err := writer.Write([]byte("You're Unauthorized!"))
// 			if err != nil {
// 				return nil, err

// 			}
// 		}
// 	})

// 	if err != nil {
// 		return err
// 	}

// }
