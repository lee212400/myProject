package db

import (
	"database/sql"

	"github.com/lee212400/myProject/domain/entity"
)

type Db interface {
	GetDbClient(ctx *entity.Context) (*sql.Tx, error)
}
