package validator

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"net/smtp"

	"github.com/aouiniamine/whatsup/backend/internal/organisms/db"
	"github.com/aouiniamine/whatsup/backend/internal/organisms/structs"
	"github.com/golang-jwt/jwt/v5"
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

var secretKey []byte = []byte("____________MY_Secret_Key____________")

func VerifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id, ok := claims["id"].(string)
		if !ok {
			fmt.Println()
			return "", errors.New("ID claim not found in JWT")
		}
		return id, nil

	} else {
		return "", errors.New("token is invalid")
	}
}

func CreateToken(id int) (string, error) {
	claims := &jwt.MapClaims{
		"id": fmt.Sprintf("%d", id),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	tokenString, err := token.SignedString(secretKey)
	return tokenString, err
}
