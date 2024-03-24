package validator

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/smtp"

	"github.com/aouiniamine/whatsup/backend/internal/organisms/db"
	"github.com/aouiniamine/whatsup/backend/internal/organisms/structs"
)

func randRangeCrypto() (int64, error) {
	max := int64(999999)
	min := int64(100001)

	diff := big.NewInt(max - min)

	n, err := rand.Int(rand.Reader, diff)
	if err != nil {
		return 0, err
	}
	return n.Int64() + min, nil
}

func ValidateWithEmail(user structs.User) error {
	generatedInt, err := randRangeCrypto()
	if err != nil {
		return err
	}
	fmt.Println("validation int:", generatedInt)
	var validator int

	db := db.DBConnection

	if err := db.QueryRow(
		"INSERT INTO connection_req (validator, user_id, req_time) "+
			"VALUES ($1, $2, CURRENT_TIMESTAMP) RETURNING validator",

		generatedInt, user.Id).Scan(&validator); err != nil {
		fmt.Println("database error on connection request creation",
			err.Error())

		return err
	}

	if err := sendValidationEmail(user.Email, validator); err != nil {
		return err
	}

	return nil
}

func sendValidationEmail(receiver string, validationCode int) error {

	content := fmt.Sprintf("Your validation code is: %d", validationCode)

	sender := fmt.Sprintf("From: <%s>\r\n", From)
	to := fmt.Sprintf("To: <%s>\r\n", receiver)
	subject := "Subject: " + content + "\r\n"
	body := content + "\r\nBye\r\n"

	msg := sender + to + subject + "\r\n" + body
	fmt.Println(msg)
	messageFormat := []byte(msg)

	// Authentication.
	auth := smtp.PlainAuth("", From, Password, SmtpHost)

	// Sending email.
	err := smtp.SendMail(SmtpHost+":"+SmtpPort, auth, From, []string{receiver}, messageFormat)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Email Sent Successfully!")
	return nil
}

var SecretKey []byte = []byte("____________MY_Secret_Key____________")
