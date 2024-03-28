package middlewares

import (
	"log"
	"net/http"

	handleErrors "github.com/aouiniamine/whatsup/backend/internal/organisms/errors"
	"github.com/aouiniamine/whatsup/backend/internal/organisms/validator"
)

func Authorize(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		id, err := validator.VerifyToken(token)
		if err != nil {
			log.Println(err)
			handleErrors.Unauthorized(w)
			return
		}

		r.Header.Set("UserId", id)
		next.ServeHTTP(w, r)
	}
}
