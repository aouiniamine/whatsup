package connections

import "time"

type ConnectionReq struct {
	Id   int       `db:"user_id"`
	Time time.Time `db:"time"`
}
