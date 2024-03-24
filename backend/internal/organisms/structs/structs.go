package structs

import "time"

// fix for cross import error
type User struct {
	Id    int    `db:"id"`
	Name  string `db:"username" json:"username"`
	Email string `db:"email" json:"email"`
}

type ConnectionReq struct {
	Id   int       `db:"user_id"`
	Time time.Time `db:"time"`
}
