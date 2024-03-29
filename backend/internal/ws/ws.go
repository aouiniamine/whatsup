package ws

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aouiniamine/whatsup/backend/internal/organisms/db"
	"github.com/aouiniamine/whatsup/backend/internal/organisms/errors"
	"github.com/aouiniamine/whatsup/backend/internal/organisms/validator"
	"github.com/aouiniamine/whatsup/backend/internal/ws/messages"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// type connection struct {
// 	Id   string
// 	Conn websocket.Conn
// }

// var connMap = make(map[string]connection)

func handleWS(username, userId string, conn *websocket.Conn) {

	AddConn(username, userId, *conn)

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			if messageType != websocket.CloseMessage {
				log.Println("Error reading message:", err)
			}
			break // Disconnect or other error
		}

		if messageType == websocket.CloseMessage {
			log.Println("Client disconnected")
			break // Client closed the connection
		}
		messages.HandleMessages(username, userId, message, *conn)
	}
	RmConn(username)
	conn.Close()

}

func OnConnect(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	userId := r.Header.Get("UserId")

	conn, err := upgrader.Upgrade(w, r, nil)
	if (err) != nil {
		errors.InternalServerError(w)
		return
	}

	go handleWS(username, userId, conn)
	// connMap[username] = connection{Id: userId, Conn: conn}
	log.Print("user: ", username, ", id:", userId, " is connected")
}

func ValidateWS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.URL.Query().Get("token")
		fmt.Println("my token: ", token)
		username := mux.Vars(r)["username"]
		id, err := validator.VerifyToken(token)
		if err != nil {
			log.Println(err)
			errors.Unauthorized(w)
			return
		}
		fmt.Println("myId", id)
		var validName string
		db := db.DBConnection
		if err := db.QueryRow("SELECT username FROM users WHERE id = $1",
			id).Scan(&validName); err != nil {
			log.Println(err)
			errors.InternalServerError(w)
			return
		}
		if username != validName {
			log.Panicln(username, " is unauthorized")
			errors.Unauthorized(w)
			return
		}

		r.Header.Set("UserId", id)

		next.ServeHTTP(w, r)
	}
}
