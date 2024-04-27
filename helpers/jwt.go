package helpers

import (
	"errors"
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

func SignToken(id uint) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
	})

	tokenString, err := token.SignedString([]byte("SECRET"))

	if err != nil {
		log.Fatal(err)
	}

	return tokenString
}

func VerifyToken(token string) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte("SECRET"), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok || !parsedToken.Valid {
		return nil, errors.New("JWT ERROR")
	}

	return claims, nil
}
