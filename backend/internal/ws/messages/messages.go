package messages

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type UserMessage struct {
	ToName  string `json:"username"`
	Message string `json:"message"`
}

func HandleMessages(username, userId string, messageByte []byte, conn websocket.Conn) {
	var message UserMessage
	if err := json.Unmarshal(messageByte, &message); err != nil {

		return
	}
	log.Println(username, " is sending to: ", message.ToName, ": ", message.Message)
}
