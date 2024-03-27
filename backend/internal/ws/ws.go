package ws

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func OnConnect(w http.ResponseWriter, r *http.Request) {

	_, err := upgrader.Upgrade(w, r, nil)
	if (err) != nil {
		panic(err)
	}
	log.Print("user is connected")
}
