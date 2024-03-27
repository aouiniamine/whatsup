package middlewares

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	handleErrors "github.com/aouiniamine/whatsup/backend/internal/organisms/errors"
	"github.com/aouiniamine/whatsup/backend/internal/organisms/validator"
	"github.com/golang-jwt/jwt/v5"
)

func verifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return validator.SecretKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id, ok := claims["id"].(string)
		if !ok {
			fmt.Println()
			return "", errors.New("ID claim not found in JWT")
		}
		return id, nil

	} else {
		return "", errors.New("token is invalid")
	}
}

func CreateToken(id int) (string, error) {
	claims := &jwt.MapClaims{
		"id": fmt.Sprintf("%d", id),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	tokenString, err := token.SignedString(validator.SecretKey)
	return tokenString, err
}

func Authorize(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		id, err := verifyToken(token)
		if err != nil {
			log.Println(err)
			handleErrors.Unauthorized(w)
			return
		}

		r.Header.Set("UserId", id)
		next.ServeHTTP(w, r)
	}
}
