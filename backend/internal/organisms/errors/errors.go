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
