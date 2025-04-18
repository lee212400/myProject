package usecase

import (
	"github.com/lee212400/myProject/domain/entity"
	"github.com/lee212400/myProject/usecase/dto"
)

type UserOutputPort interface {
	GetUser(ctx *entity.Context, out *dto.GetUserOutputDto) error
	CreateUser(ctx *entity.Context, out *dto.CreateUserOutputDto) error
	UpdateUser(ctx *entity.Context, out *dto.UpdateUserOutputDto) error
	DeleteUser(ctx *entity.Context, out *dto.DeleteUserOutputDto) error
}
