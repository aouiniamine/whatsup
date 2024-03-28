package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/aouiniamine/whatsup/backend/internal/organisms/db"
	"github.com/aouiniamine/whatsup/backend/internal/organisms/errors"
	"github.com/aouiniamine/whatsup/backend/internal/organisms/structs"
	"github.com/aouiniamine/whatsup/backend/internal/organisms/validator"
	"github.com/aouiniamine/whatsup/backend/internal/repositories/user"
)

type ConnectBody struct {
	Credential string `json:"credential"`
}

type ValidateReq struct {
	Email string `json:"email" db:"email"`
	Code  int    `json:"code" db:"validator"`
}

type ValidateRes struct {
	Email string `json:"email"`
}

type User struct {
	Email string `json:"email"`
	Name  string `json:"username"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var body User
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		errors.InternalServerError(w)

		return
	}
	user, err := user.AddUser(body.Email, body.Name)
	if err != nil {
		errors.UserAlreadyExist(w, err)
		return
	}
	fmt.Println("user is created")
	if err := validator.ValidateWithEmail(user); err != nil {
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
	user, err := user.GetByCredential(body.Credential)
	if err != nil {
		errors.InternalServerError(w)
		fmt.Println(err)
		return
	}
	if err := validator.ValidateWithEmail(user); err != nil {
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
	id := r.Header.Get("UserId")
	userId, err := strconv.Atoi(id)
	if err != nil {
		errors.InternalServerError(w)
		return
	}
	user, err := user.GetById(userId)
	if err != nil {
		errors.InternalServerError(w)
		return
	}

	userRes := User{Name: user.Name, Email: user.Email}
	jsonRes, err := json.Marshal(userRes)
	if err != nil {
		errors.InternalServerError(w)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonRes)

}
