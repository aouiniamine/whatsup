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
	Id    int    `db:"id" json:"id"`
	Name  string `db:"username" json:"username"`
	Email string `db:"email" json:"email"`
}

func AuthRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/auth/connect", connect).Methods("POST")
	return r
}

func connect(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request has came!!")
	var body ConnectBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		errors.InternalServerError(w, r)
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
		errors.InternalServerError(w, r)
		return
	}

	jsonRes, err := json.Marshal(user)
	if err != nil {
		errors.InternalServerError(w, r)
		return
	}

	w.Write(jsonRes)
}
