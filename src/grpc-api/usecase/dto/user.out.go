package dto

import "github.com/lee212400/myProject/domain/entity"

type GetUserOutputDto struct {
	User *entity.User
}

type CreateUserOutputDto struct {
	User *entity.User
}

type UpdateUserOutputDto struct {
	User *entity.User
}

type DeleteUserOutputDto struct{}
