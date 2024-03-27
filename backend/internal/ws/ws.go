package ws

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func OnConnect(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]

	_, err := upgrader.Upgrade(w, r, nil)
	if (err) != nil {
		panic(err)
	}
	log.Print("user: ", username, " is connected")
}
