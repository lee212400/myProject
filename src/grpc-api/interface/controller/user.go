package controller

import (
	"github.com/lee212400/myProject/domain/entity"
	"github.com/lee212400/myProject/usecase"
	"github.com/lee212400/myProject/usecase/dto"

	pb "github.com/lee212400/myProject/rpc/proto"
)

type UserController struct {
	inputPort usecase.UserInputPort
}

func NewUserController(inputPort usecase.UserInputPort) *UserController {
	return &UserController{
		inputPort: inputPort,
	}
}

func (i *UserController) GetUser(ctx *entity.Context, in *pb.GetUserRequest) error {
	dto := &dto.GetUserInputDto{
		UserId: in.UserId,
	}
	return i.inputPort.GetUser(ctx, dto)
}
func (i *UserController) CreateUser(ctx *entity.Context, in *pb.CreateUserRequest) error {
	dto := &dto.CreateUserInputDto{
		User: &entity.User{
			FirstName: in.User.FirstName,
			LastName:  in.User.LastName,
			Email:     in.User.Email,
			Age:       in.User.Age,
		},
	}
	return i.inputPort.CreateUser(ctx, dto)
}
func (i *UserController) UpdateUser(ctx *entity.Context, in *pb.UpdateUserRequest) error {
	dto := &dto.UpdateUserInputDto{
		UserId: in.UserId,
		Age:    in.Age,
	}
	return i.inputPort.UpdateUser(ctx, dto)
}
func (i *UserController) DeleteUser(ctx *entity.Context, in *pb.DeleteUserRequest) error {
	dto := &dto.DeleteUserInputDto{
		UserId: in.UserId,
	}
	return i.inputPort.DeleteUser(ctx, dto)
}
