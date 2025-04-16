package dto

import "github.com/lee212400/myProject/domain/entity"

type GetUserInputDto struct {
	UserId string
}

type CreateUserInputDto struct {
	User *entity.User
}

type UpdateUserInputDto struct {
	UserId string
	Age    int32
}

type DeleteUserInputDto struct {
	UserId string
}
