package errors

import "net/http"

func InternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 Internal Server Error"))
}

func UserAlreadyExist(w http.ResponseWriter, errorMsg error) {
	w.WriteHeader(http.StatusConflict)
	w.Write([]byte("User already exists: 409"))
}

func SessionHasExpired(w http.ResponseWriter) {
	w.WriteHeader(440)
	w.Write([]byte("Session Has Expired!"))
}

func Unauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("User is unauthorized!"))
}
