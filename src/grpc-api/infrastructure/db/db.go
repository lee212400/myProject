package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lee212400/myProject/domain/entity"
)

func GetDb(ctx *entity.Context, db *sql.DB) (*sql.Tx, error) {
	if val, ok := ctx.Session["mysql"]; ok {
		if tx, ok := val.(*sql.Tx); ok {
			return tx, nil
		}
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	ctx.Session["mysql"] = tx

	return tx, nil

}

func CloseDb(ctx *entity.Context, sucess bool) {
	if val, ok := ctx.Session["mysql"]; ok {
		if tx, ok := val.(*sql.Tx); ok {
			if sucess {
				log.Println("commit")
				_ = tx.Commit()
			} else {
				log.Println("rollback")
				_ = tx.Rollback()
			}
			delete(ctx.Session, "mysql")
		} else {

		}
	}
}

func NewDb() *sql.DB {
	dsn := "root:password@tcp(mysql.default.svc.cluster.local:3306)/mydb?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to DB: %v", err))
	}
	return db
}
