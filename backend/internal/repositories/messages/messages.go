package messages

import (
	"log"

	"github.com/aouiniamine/whatsup/backend/internal/organisms/db"
)

type Message struct {
	Id      int    `db:"id" json:"id"`
	Message string `db:"message_content" json:"message"`
	Name    string `db:"username" json:"username"`
}

func CreateMessage(from, to int, message string) error {
	db := db.DBConnection
	var id int
	if err := db.QueryRow(
		"INSERT INTO messages (sender, recipient, message_content)"+
			" VALUES ($1, $2, $3) RETURNING id", from, to, message).Scan(&id); err != nil {
		return err
	}
	log.Println("message is created!")
	return nil
}

func getAllUserMessages(userId int) ([]Message, error) {
	db := db.DBConnection
	var allMessage []Message
	rows, err := db.Query("SELECT messages.id AS id, users.username AS username, messages.message_content AS message FROM messages "+
		"INNER JOIN users ON messages.sender = users.id WHERE messages.sender = $1 OR recipient = $1", userId)
	if err != nil {
		return allMessage, err
	}
	if err := rows.Scan(allMessage); err != nil {
		return allMessage, err
	}
	return allMessage, nil

}
