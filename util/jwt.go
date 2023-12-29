package util

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtAdapter struct {
}

func (j *JwtAdapter) CreateToken(email string, username string) (string, error) {
	mapClaims := jwt.MapClaims{
		"username": username,
		"email":    email,
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
