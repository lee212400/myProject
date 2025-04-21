package usecase

import (
	"github.com/lee212400/myProject/domain/entity"
	"github.com/lee212400/myProject/usecase/dto"
)

type UserInputPort interface {
	GetUser(ctx *entity.Context, in *dto.GetUserInputDto) error
	CreateUser(ctx *entity.Context, in *dto.CreateUserInputDto) error
	UpdateUser(ctx *entity.Context, in *dto.UpdateUserInputDto) error
	DeleteUser(ctx *entity.Context, in *dto.DeleteUserInputDto) error
}
