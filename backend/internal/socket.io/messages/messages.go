package messages

import (
	"fmt"

	socketio "github.com/googollee/go-socket.io"
)

func OnConnect(s socketio.Conn) error {

	s.SetContext("")
	fmt.Println("User:", s.ID(), "is connected to ws")
	return nil
}

func OnError(s socketio.Conn, err error) {
	fmt.Println("socket error", err)
}
