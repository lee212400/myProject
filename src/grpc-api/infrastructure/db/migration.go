package db

import (
	"database/sql"
	"fmt"
	"log"
)

func Migration() {
	db := getDb()

	existTb := tableExists(db, "users")

	if existTb {
		return
	}

	query := `CREATE TABLE users (
    user_id VARCHAR(20) NOT NULL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
	age INT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);`

	_, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
}

func tableExists(db *sql.DB, tableName string) bool {
	var exists string
	query := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	err := db.QueryRow(query).Scan(&exists)

	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		log.Fatal(err)
	}

	return true
}

func getDb() *sql.DB {
	dsn := "root:password@tcp(localhost:3306)/mydb?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to DB: %v", err))
	}
	return db
}
