package messages

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/aouiniamine/whatsup/backend/internal/repositories/messages"
	"github.com/aouiniamine/whatsup/backend/internal/repositories/user"
	"github.com/aouiniamine/whatsup/backend/internal/ws/connStore"
	"github.com/gorilla/websocket"
)

type UserMessage struct {
	ToName  string `json:"username"`
	Message string `json:"message"`
}

type SendMessage struct {
	From    string `json:"username"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

func HandleMessages(username string, userId int, messageByte []byte, conn *websocket.Conn) {
	var message UserMessage
	if err := json.Unmarshal(messageByte, &message); err != nil {

		return
	}
	var to int
	recipientWSConn, ok := connStore.GetConn(message.ToName)
	if !ok {
		recipient, err := user.GetByCredential(message.ToName)
		if err != nil {
			if err := conn.WriteControl(websocket.CloseMessage, []byte("user not found"), time.Now()); err != nil {
				fmt.Println("couldn't notify user", err)
				return
			}
			fmt.Println("user not found")
			return
		}
		to = recipient.Id
	}
	if err := messages.CreateMessage(userId, to, message.Message); err != nil {
		println("create messsage error:", err.Error())
		return
	}
	if ok {
		to = recipientWSConn.Id
		recipientWSConn.Conn.WriteJSON(SendMessage{From: username, Message: message.Message, Type: "message:recieve"})
	}
	log.Println(username, " is sending to: ", message.ToName, ", ", to, ": ", message.Message)
}
