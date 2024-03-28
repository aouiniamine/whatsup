package auth

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aouiniamine/whatsup/backend/internal/organisms/db"
	"github.com/aouiniamine/whatsup/backend/internal/organisms/errors"
	"github.com/aouiniamine/whatsup/backend/internal/organisms/structs"
	"github.com/aouiniamine/whatsup/backend/internal/organisms/validator"
)

type ConnectBody struct {
	Credential string `json:"credential"`
}

type ValidateReq struct {
	Email string `json:"email" db:"email"`
	Code  int    `json:"code" db:"validator"`
}

// type ValidteRes struct {
// 	Token string `json: token`
// }

type ValidateRes struct {
	Email string `json:"email"`
}

type GetUserRes struct {
	Email string `json:"email"`
	Name  string `json:"username"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var body structs.User
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		errors.InternalServerError(w)

		return
	}
	db := db.DBConnection

	if err := db.QueryRow(
		"INSERT INTO users (email, username) VALUES ($1, $2) RETURNING id",
		body.Email, body.Name).Scan(&body.Id); err != nil {
		fmt.Println("database error on register", err.Error())
		errors.UserAlreadyExist(w, err)
		return
	}
	fmt.Println("user is created")
	err := validator.ValidateWithEmail(body)
	if err != nil {
		fmt.Println("error on email validation:", err)
		errors.InternalServerError(w)
		return
	}

	jsonRes, err := json.Marshal(ValidateRes{Email: body.Email})
	if err != nil {
		errors.InternalServerError(w)
		return
	}

	w.Write(jsonRes)

}

func Connect(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request has came!!")
	var body ConnectBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		errors.InternalServerError(w)

		return
	}
	db := db.DBConnection
	var user structs.User
	row := db.QueryRow("SELECT * FROM users WHERE username = $1 OR email = $1", body.Credential)

	if err := row.Scan(&user.Id, &user.Name, &user.Email); err == sql.ErrNoRows {
		w.Write([]byte("User not found!"))
		return
	} else if err != nil {
		fmt.Println(err)
		errors.InternalServerError(w)

		return
	}
	err := validator.ValidateWithEmail(user)
	if err != nil {
		fmt.Println("error on email validation:", err)
		errors.InternalServerError(w)
		return
	}

	jsonRes, err := json.Marshal(ValidateRes{Email: user.Email})
	if err != nil {
		errors.InternalServerError(w)

		return
	}

	w.Write(jsonRes)
}

func Validate(w http.ResponseWriter, r *http.Request) {
	var body ValidateReq
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		errors.InternalServerError(w)
		fmt.Println(err)
		return
	}
	db := db.DBConnection
	var connectionReq structs.ConnectionReq
	err := db.QueryRow("SELECT connection_req.user_id, connection_req.req_time "+
		"FROM connection_req INNER JOIN users ON users.id = connection_req.user_id "+
		"WHERE users.email = $1 AND connection_req.validator = $2",
		body.Email, body.Code).Scan(&connectionReq.Id, &connectionReq.Time)

	if err != nil {
		errors.InternalServerError(w)
		fmt.Println(err)
		return
	}
	timeNow := time.Now()
	experationDate := connectionReq.Time.Add(2 * time.Hour)
	fmt.Println("is expired: ", experationDate.Before(timeNow))
	sessionHasExpired := experationDate.Before(timeNow)
	if sessionHasExpired {
		errors.SessionHasExpired(w)
		return

	}

	token, err := validator.CreateToken(connectionReq.Id)
	if err != nil {
		errors.InternalServerError(w)
		return
	}
	fmt.Println("token:", token)
	w.Write([]byte(token))

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("UserId")
	fmt.Println(userId)
	user := GetUserRes{}
	db := db.DBConnection
	if err := db.QueryRow("SELECT username, email FROM users WHERE id = $1",
		userId).Scan(&user.Name, &user.Email); err != nil {
		errors.InternalServerError(w)
	}
	jsonRes, err := json.Marshal(user)
	if err != nil {
		errors.InternalServerError(w)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonRes)

}
