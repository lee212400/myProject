package db

import (
	"database/sql"
	"errors"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lee212400/myProject/domain/entity"
)

type MockDb struct {
	Db   *sql.DB
	Mock sqlmock.Sqlmock
}

func NewMockDb() *MockDb {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic(err)
	}

	return &MockDb{
		Db:   db,
		Mock: mock,
	}
}

func (m *MockDb) GetDbClient(ctx *entity.Context) (*sql.Tx, error) {
	return m.Db.Begin()
}

func (m *MockDb) SelectError(mock sqlmock.Sqlmock) {
	mock.ExpectQuery("select * from").WillReturnError(errors.New("failed to select"))
}

func (m *MockDb) CreateError(mock sqlmock.Sqlmock) {
	mock.ExpectQuery("insert into").WillReturnError(errors.New("failed to insert int"))
}

func (m *MockDb) UpdateError(mock sqlmock.Sqlmock) {
	mock.ExpectQuery("update").WillReturnError(errors.New("failed to update"))
}

func (m *MockDb) DeleteError(mock sqlmock.Sqlmock) {
	mock.ExpectQuery("delete from").WillReturnError(errors.New("failed to delete"))
}
