package db

import (
	"database/sql"
	"fmt"
)

func NewDb() *sql.DB {
	dsn := "user:password@tcp(mysql:3306)/mydb?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to DB: %v", err))
	}
	return db
}
