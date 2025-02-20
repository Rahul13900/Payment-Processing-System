package middleware

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func CreateConnection(dbURL string) *sql.DB {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully Connected to Postgres")
	return db
}
