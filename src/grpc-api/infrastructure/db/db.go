package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func NewDb() *sql.DB {
	dsn := "root:password@tcp(mysql:3306)/mydb?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to DB: %v", err))
	}
	return db
}
