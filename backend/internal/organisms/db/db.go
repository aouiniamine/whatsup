package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func connection() sqlx.DB {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=whatsup sslmode=disable password=123456 host=localhost")
	if err != nil {
		log.Fatalln(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("DB: Connected Successfuly")
	}

	return *db

}

var DBConnection sqlx.DB = connection()
