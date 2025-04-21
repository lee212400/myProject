package repository

import "github.com/lee212400/myProject/domain/entity"

type UserRepository interface {
	GetUser(ctx *entity.Context, userId string) (*entity.User, error)
	CreateUser(ctx *entity.Context, firstName string, lastName string, email string, age int32) (string, error)
	UpdateUser(ctx *entity.Context, userId string, age int32) error
	DeleteUser(ctx *entity.Context, userId string) error
}
