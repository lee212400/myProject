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
	return nil
}
func (i *UserController) UpdateUser(ctx *entity.Context, in *pb.UpdateUserRequest) error {
	return nil
}
func (i *UserController) DeleteUser(ctx *entity.Context, in *pb.DeleteUserRequest) error {
	return nil
}
