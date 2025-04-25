package controller

import (
	"github.com/lee212400/myProject/domain/entity"
	"github.com/lee212400/myProject/usecase"
	"github.com/lee212400/myProject/usecase/dto"

	pb "github.com/lee212400/myProject/rpc/proto"
	ue "github.com/lee212400/myProject/utils/errors"
	log "github.com/lee212400/myProject/utils/logger"
	validate "github.com/lee212400/myProject/utils/validate"
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
	log.WithContext(ctx).Debug("start GetUser")
	// if err := validate.Validate.Validate(in); err != nil {
	// 	return ue.WithError(ctx, ue.InvalidArgument, err)
	// }

	dto := &dto.GetUserInputDto{
		UserId: in.UserId,
	}
	return i.inputPort.GetUser(ctx, dto)
}
func (i *UserController) CreateUser(ctx *entity.Context, in *pb.CreateUserRequest) error {
	if err := validate.Validate.Validate(in); err != nil {
		return ue.New(ctx, ue.InvalidArgument, "Invaild request data")
	}

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
	if err := validate.Validate.Validate(in); err != nil {
		return ue.New(ctx, ue.InvalidArgument, "Invaild request data")
	}

	dto := &dto.UpdateUserInputDto{
		UserId: in.UserId,
		Age:    in.Age,
	}
	return i.inputPort.UpdateUser(ctx, dto)
}
func (i *UserController) DeleteUser(ctx *entity.Context, in *pb.DeleteUserRequest) error {
	if err := validate.Validate.Validate(in); err != nil {
		return ue.New(ctx, ue.InvalidArgument, "Invaild request data")
	}

	dto := &dto.DeleteUserInputDto{
		UserId: in.UserId,
	}
	return i.inputPort.DeleteUser(ctx, dto)
}
