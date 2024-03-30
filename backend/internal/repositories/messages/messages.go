package messages

import (
	"log"

	"github.com/aouiniamine/whatsup/backend/internal/organisms/db"
)

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
