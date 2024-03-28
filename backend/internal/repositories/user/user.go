package user

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/aouiniamine/whatsup/backend/internal/organisms/db"
)

type User struct {
	Id    int    `db:"id"`
	Name  string `db:"username" json:"username"`
	Email string `db:"email" json:"email"`
}

func GetById(id int) (User, error) {
	var user User
	db := db.DBConnection
	if err := db.QueryRow("SELECT * FROM users WHERE id = $1",
		id).Scan(&user.Id, &user.Name, &user.Email); err != nil {
		return user, err
	}
	return user, nil
}

func AddUser(email, username string) (User, error) {
	var user User
	db := db.DBConnection

	if err := db.QueryRow(
		"INSERT INTO users (email, username) VALUES ($1, $2) RETURNING *",
		email, username).Scan(&user.Id, &user.Name, &user.Email); err != nil {
		fmt.Println("database error on register", err.Error())

		return user, err
	}
	return user, nil
}

func GetByCredential(credential string) (User, error) {
	var user User
	db := db.DBConnection
	row := db.QueryRow("SELECT * FROM users WHERE username = $1 OR email = $1", credential)

	if err := row.Scan(&user.Id, &user.Name, &user.Email); err == sql.ErrNoRows {
		// w.Write()
		return user, errors.New("User not found")
	} else if err != nil {
		return user, err
	}
	return user, nil
}
