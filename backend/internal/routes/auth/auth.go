package auth

import (
	"net/http"

	"github.com/gorilla/mux"
)

func AuthRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/auth/signup", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("api works"))
	}).Methods("POST")
	return r
}
