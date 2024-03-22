package auth

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aouiniamine/whatsup/backend/internal/organisms/db"
	"github.com/aouiniamine/whatsup/backend/internal/organisms/errors"
	"github.com/gorilla/mux"
)

type ConnectBody struct {
	Credential string `json:"credential"`
}

type User struct {
	Id    int    `db:"id"`
	Name  string `db:"username" json:"username"`
	Email string `db:"email" json:"email"`
}

func AuthRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/auth/connect", connect).Methods("POST")
	r.HandleFunc("/auth/register", register).Methods("POST")
	return r
}

func register(w http.ResponseWriter, r *http.Request) {
	var body User
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		errors.InternalServerError(w)

		return
	}
	db := db.DBConnection

	if err := db.QueryRow(
		"INSERT INTO users (email, username) VALUES ($1, $2) RETURNING email",
		body.Email, body.Name).Scan(&body.Email); err != nil {
		fmt.Println("database error on register", err.Error())
		errors.UserAlreadyExist(w, err)
		return
	}

	jsonRes, err := json.Marshal(body)
	if err != nil {
		errors.InternalServerError(w)
		return
	}

	w.Write(jsonRes)

}

func connect(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request has came!!")
	var body ConnectBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		errors.InternalServerError(w)

		return
	}
	db := db.DBConnection
	var user User
	row := db.QueryRow("SELECT * FROM users WHERE username = $1 OR email = $2", body.Credential, body.Credential)

	if err := row.Scan(&user.Id, &user.Email, &user.Name); err == sql.ErrNoRows {
		w.Write([]byte("User not found!"))
		return
	} else if err != nil {
		fmt.Println(err)
		errors.InternalServerError(w)

		return
	}

	jsonRes, err := json.Marshal(user)
	if err != nil {
		errors.InternalServerError(w)

		return
	}

	w.Write(jsonRes)
}
