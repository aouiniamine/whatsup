package main

import (
	"fmt"
	"net/http"

	"github.com/aouiniamine/whatsup/backend/internal/routes/auth"
)

func main() {
	router := auth.AuthRouter()

	port := 8080
	addr := fmt.Sprintf("192.168.0.4:%d", port)
	fmt.Println("Server running on port:", port)
	fmt.Println(addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}

}
