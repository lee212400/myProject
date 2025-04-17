package repository

import (
	"database/sql"
	"fmt"

	"github.com/lee212400/myProject/domain/entity"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepositoryImpl(db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (i *UserRepositoryImpl) GetUser(ctx *entity.Context, userId string) (*entity.User, error) {
	fmt.Println("Repository GetUser")
	return &entity.User{}, nil
}
func (i *UserRepositoryImpl) CreateUser(ctx *entity.Context, firstName string, lastName string, email string, age int32) (*entity.User, error) {
	return &entity.User{}, nil
}
func (i *UserRepositoryImpl) UpdateUser(ctx *entity.Context, userId string, age int32) (*entity.User, error) {
	return &entity.User{}, nil
}
func (i *UserRepositoryImpl) DeleteUser(ctx *entity.Context, userId string) error { return nil }
