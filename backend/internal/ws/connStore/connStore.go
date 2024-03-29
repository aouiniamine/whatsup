package ws

import (
	"sync"

	"github.com/gorilla/websocket"
)

type WSConn struct {
	Id   string
	Conn websocket.Conn
}

// var connectedUsers = map[string]WSConn

var AllConns sync.Map

func AddConn(username string, userId string, conn websocket.Conn) {
	AllConns.Store(username, WSConn{Id: userId, Conn: conn})

}

func RmConn(username string) {
	AllConns.Delete(username)
}
