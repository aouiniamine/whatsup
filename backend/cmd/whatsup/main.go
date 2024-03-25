package main

import (
	"fmt"
	"net/http"

	"github.com/aouiniamine/whatsup/backend/internal/routes/auth"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Auth API's
	user := router.PathPrefix("/api/auth").Subrouter()
	user.HandleFunc("/connect", auth.Connect).Methods("POST")
	user.HandleFunc("/register", auth.Register).Methods("POST")
	user.HandleFunc("/validate", auth.Validate).Methods("POST")

	port := 8080
	addr := fmt.Sprintf("192.168.0.14:%d", port)
	fmt.Println("Server running on port:", port)
	fmt.Println(addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}

}
