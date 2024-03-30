package connStore

import (
	"sync"

	"github.com/gorilla/websocket"
)

type WSConn struct {
	Id   int
	Conn websocket.Conn
}

var connectedUsers = make(map[string]WSConn)

var AllConns sync.Map

func AddConn(username string, userId int, conn websocket.Conn) {
	// AllConns.Store(username, WSConn{Id: userId, Conn: conn})
	connectedUsers[username] = WSConn{Id: userId, Conn: conn}

}

func RmConn(username string) {
	delete(connectedUsers, username)
}

func GetConn(username string) (WSConn, bool) {
	wsConn, ok := connectedUsers[username]
	return wsConn, ok
}
