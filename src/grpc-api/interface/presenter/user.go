package presenter

import (
	"github.com/lee212400/myProject/domain/entity"
	pb "github.com/lee212400/myProject/rpc/proto"
	"github.com/lee212400/myProject/usecase/dto"
)

type UserPresenter struct {
}

func NewUserPresenter() *UserPresenter {
	return &UserPresenter{}
}

func (i *UserPresenter) GetUser(ctx *entity.Context, out *dto.GetUserOutputDto) error {
	ctx.Response = &pb.GetUserResponse{
		User: &pb.User{
			UserId:    out.User.UserId,
			Email:     out.User.Email,
			FirstName: out.User.FirstName,
			LastName:  out.User.LastName,
			Age:       out.User.Age,
		},
	}
	return nil
}
func (i *UserPresenter) CreateUser(ctx *entity.Context, out *dto.CreateUserOutputDto) error {
	ctx.Response = &pb.CreateUserResponse{
		User: &pb.User{
			UserId:    out.User.UserId,
			Email:     out.User.Email,
			FirstName: out.User.FirstName,
			LastName:  out.User.LastName,
			Age:       out.User.Age,
		},
	}
	return nil
}
func (i *UserPresenter) UpdateUser(ctx *entity.Context, out *dto.UpdateUserOutputDto) error {
	ctx.Response = &pb.UpdateUserResponse{
		User: &pb.User{
			UserId:    out.User.UserId,
			Email:     out.User.Email,
			FirstName: out.User.FirstName,
			LastName:  out.User.LastName,
			Age:       out.User.Age,
		},
	}
	return nil
}
func (i *UserPresenter) DeleteUser(ctx *entity.Context, out *dto.DeleteUserOutputDto) error {
	ctx.Response = &pb.DeleteUserResponse{}
	return nil
}
