package presenter

import (
	"context"
	"testing"

	"github.com/lee212400/myProject/domain/entity"
	pb "github.com/lee212400/myProject/rpc/proto"
	"github.com/lee212400/myProject/usecase/dto"
	uc "github.com/lee212400/myProject/utils/context"
	"github.com/stretchr/testify/require"
)

func TestUserPresenter_GetUser(t *testing.T) {
	ct := uc.NewContext(context.Background())

	presenter := NewUserPresenter()

	type args struct {
		ctx *entity.Context
		out *dto.GetUserOutputDto
	}
	tests := []struct {
		name    string
		args    args
		wantDt  *pb.GetUserResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"correct",
			args{ct, &dto.GetUserOutputDto{
				User: &entity.User{
					UserId:    "123412wq12q123q123q1",
					Email:     "test@test.com",
					FirstName: "first",
					LastName:  "last",
					Age:       30,
				},
			}},
			&pb.GetUserResponse{
				User: &pb.User{
					UserId:    "123412wq12q123q123q1",
					Email:     "test@test.com",
					FirstName: "first",
					LastName:  "last",
					Age:       30,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := presenter.GetUser(tt.args.ctx, tt.args.out); (err != nil) != tt.wantErr {
				t.Errorf("UserPresenter.GetUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			require.Equal(t, tt.wantDt, tt.args.ctx.Response.(*pb.GetUserResponse))
		})
	}
}

func TestUserPresenter_CreateUser(t *testing.T) {
	ct := uc.NewContext(context.Background())

	presenter := NewUserPresenter()

	type args struct {
		ctx *entity.Context
		out *dto.CreateUserOutputDto
	}
	tests := []struct {
		name    string
		args    args
		wantDt  *pb.CreateUserResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"correct",
			args{ct, &dto.CreateUserOutputDto{
				User: &entity.User{
					UserId:    "123412wq12q123q123q1",
					Email:     "test@test.com",
					FirstName: "first",
					LastName:  "last",
					Age:       30,
				},
			}},
			&pb.CreateUserResponse{
				User: &pb.User{
					UserId:    "123412wq12q123q123q1",
					Email:     "test@test.com",
					FirstName: "first",
					LastName:  "last",
					Age:       30,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := presenter.CreateUser(tt.args.ctx, tt.args.out); (err != nil) != tt.wantErr {
				t.Errorf("UserPresenter.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}

			require.Equal(t, tt.wantDt, tt.args.ctx.Response.(*pb.CreateUserResponse))
		})
	}
}

func TestUserPresenter_UpdateUser(t *testing.T) {
	ct := uc.NewContext(context.Background())

	presenter := NewUserPresenter()
	type args struct {
		ctx *entity.Context
		out *dto.UpdateUserOutputDto
	}
	tests := []struct {
		name    string
		args    args
		wantDt  *pb.UpdateUserResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"correct",
			args{ct, &dto.UpdateUserOutputDto{
				User: &entity.User{
					UserId:    "123412wq12q123q123q1",
					Email:     "test@test.com",
					FirstName: "first",
					LastName:  "last",
					Age:       30,
				},
			}},
			&pb.UpdateUserResponse{
				User: &pb.User{
					UserId:    "123412wq12q123q123q1",
					Email:     "test@test.com",
					FirstName: "first",
					LastName:  "last",
					Age:       30,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := presenter.UpdateUser(tt.args.ctx, tt.args.out); (err != nil) != tt.wantErr {
				t.Errorf("UserPresenter.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}

			require.Equal(t, tt.wantDt, tt.args.ctx.Response.(*pb.UpdateUserResponse))
		})
	}
}

func TestUserPresenter_DeleteUser(t *testing.T) {
	ct := uc.NewContext(context.Background())

	presenter := NewUserPresenter()

	type args struct {
		ctx *entity.Context
		out *dto.DeleteUserOutputDto
	}
	tests := []struct {
		name    string
		args    args
		wantDt  *pb.DeleteUserResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"correct",
			args{ct, &dto.DeleteUserOutputDto{}},
			&pb.DeleteUserResponse{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := presenter.DeleteUser(tt.args.ctx, tt.args.out); (err != nil) != tt.wantErr {
				t.Errorf("UserPresenter.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}

			require.Equal(t, tt.wantDt, tt.args.ctx.Response.(*pb.DeleteUserResponse))
		})
	}
}
