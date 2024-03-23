package structs

// fix for cross import error
type User struct {
	Id    int    `db:"id"`
	Name  string `db:"username" json:"username"`
	Email string `db:"email" json:"email"`
}

type ConnectionReq struct {
	Id    int    `db:"id"`
	Email string `db:"email" json:"email"`
	Code  int    `db:"validator"`
}
