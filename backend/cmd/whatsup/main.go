package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aouiniamine/whatsup/backend/internal/routes/auth"
	"github.com/aouiniamine/whatsup/backend/internal/socket.io/messages"
	socketio "github.com/googollee/go-socket.io"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Auth API's
	user := router.PathPrefix("/api/auth").Subrouter()
	user.HandleFunc("/connect", auth.Connect).Methods("POST")
	user.HandleFunc("/register", auth.Register).Methods("POST")
	user.HandleFunc("/validate", auth.Validate).Methods("POST")

	socket := socketio.NewServer(nil)

	socket.OnConnect("/", messages.OnConnect)
	socket.OnError("/", messages.OnError)

	router.Handle("/socket.io/", (socket))
	go func() {
		if err := socket.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()

	port := 8080
	addr := fmt.Sprintf("192.168.0.14:%d", port)
	fmt.Println("Server running on port:", port)
	fmt.Println(addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}

}
