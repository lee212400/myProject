package controller

import (
	"context"
	"fmt"
	"testing"

	"github.com/lee212400/myProject/domain/entity"
	mu "github.com/lee212400/myProject/mock/usecase"
	pb "github.com/lee212400/myProject/rpc/proto"
	"github.com/lee212400/myProject/usecase/dto"
	uc "github.com/lee212400/myProject/utils/context"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func IfElse[T any](cond bool, trueVal, falseVal T) T {
	if cond {
		return trueVal
	}
	return falseVal
}

func TestUserController_GetUser(t *testing.T) {

	ct := uc.NewContext(context.Background())
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	inputPort := mu.NewMockUserInputPort(ctrl)
	mockCtrl := NewUserController(inputPort)

	const (
		callValidate int = iota
		callInputPort
	)

	type args struct {
		ctx *entity.Context
		in  *pb.GetUserRequest
	}
	tests := []struct {
		name       string
		args       args
		callExpect int
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			"correct",
			args{ct, &pb.GetUserRequest{
				UserId: "123412wq12q123q123q1",
			}},
			callInputPort,
			false,
		},
		{
			"validate error",
			args{ct, &pb.GetUserRequest{
				UserId: "123412wq12q123",
			}},
			callValidate,
			true,
		},
		{
			"inputport error",
			args{ct, &pb.GetUserRequest{
				UserId: "123412wq12q123q123q1",
			}},
			callInputPort,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputPort.EXPECT().GetUser(gomock.Any(), gomock.Any()).Do(func(ctx *entity.Context, in *dto.GetUserInputDto) {
				require.Equal(t, tt.args.in.UserId, in.UserId)
			}).Return(IfElse(tt.wantErr == true, fmt.Errorf("dummy error"), nil)).
				Times(IfElse(tt.callExpect == callValidate, 0, 1))

			if err := mockCtrl.GetUser(tt.args.ctx, tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("UserController.GetUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserController_CreateUser(t *testing.T) {
	ct := uc.NewContext(context.Background())
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	inputPort := mu.NewMockUserInputPort(ctrl)
	mockCtrl := NewUserController(inputPort)

	const (
		callValidate int = iota
		callInputPort
	)

	type args struct {
		ctx *entity.Context
		in  *pb.CreateUserRequest
	}
	tests := []struct {
		name       string
		args       args
		callExpect int
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			"correct",
			args{ct, &pb.CreateUserRequest{
				User: &pb.User{
					Email:     "test@test.com",
					FirstName: "first",
					LastName:  "last",
					Age:       30,
				},
			}},
			callInputPort,
			false,
		},
		{
			"validate email error",
			args{ct, &pb.CreateUserRequest{
				User: &pb.User{
					Email:     "test.com",
					FirstName: "first",
					LastName:  "last",
					Age:       30,
				},
			}},
			callValidate,
			true,
		},
		{
			"input port error",
			args{ct, &pb.CreateUserRequest{
				User: &pb.User{
					Email:     "test@test.com",
					FirstName: "first",
					LastName:  "last",
					Age:       30,
				},
			}},
			callInputPort,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputPort.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Do(func(ctx *entity.Context, in *dto.CreateUserInputDto) {
				require.Equal(t, tt.args.in.User.Email, in.User.Email)
				require.Equal(t, tt.args.in.User.FirstName, in.User.FirstName)
				require.Equal(t, tt.args.in.User.LastName, in.User.LastName)
				require.Equal(t, tt.args.in.User.Age, in.User.Age)
			}).Return(IfElse(tt.wantErr == true, fmt.Errorf("dummy error"), nil)).
				Times(IfElse(tt.callExpect == callValidate, 0, 1))

			if err := mockCtrl.CreateUser(tt.args.ctx, tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("UserController.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserController_UpdateUser(t *testing.T) {
	ct := uc.NewContext(context.Background())
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	inputPort := mu.NewMockUserInputPort(ctrl)
	mockCtrl := NewUserController(inputPort)

	const (
		callValidate int = iota
		callInputPort
	)
	type args struct {
		ctx *entity.Context
		in  *pb.UpdateUserRequest
	}
	tests := []struct {
		name       string
		args       args
		callExpect int
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			"correct",
			args{ct, &pb.UpdateUserRequest{
				UserId: "123412wq12q123q123q1",
				Age:    30,
			}},
			callInputPort,
			false,
		},
		{
			"validate user id error",
			args{ct, &pb.UpdateUserRequest{
				UserId: "123412wq12q1231",
				Age:    30,
			}},
			callValidate,
			true,
		},
		{
			"validate age error",
			args{ct, &pb.UpdateUserRequest{
				UserId: "123412wq12q123q123q1",
				Age:    0,
			}},
			callValidate,
			true,
		},
		{
			"input port error",
			args{ct, &pb.UpdateUserRequest{
				UserId: "123412wq12q123q123q1",
				Age:    30,
			}},
			callInputPort,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputPort.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Do(func(ctx *entity.Context, in *dto.UpdateUserInputDto) {
				require.Equal(t, tt.args.in.UserId, in.UserId)
				require.Equal(t, tt.args.in.Age, in.Age)
			}).Return(IfElse(tt.wantErr == true, fmt.Errorf("dummy error"), nil)).
				Times(IfElse(tt.callExpect == callValidate, 0, 1))

			if err := mockCtrl.UpdateUser(tt.args.ctx, tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("UserController.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserController_DeleteUser(t *testing.T) {
	ct := uc.NewContext(context.Background())
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	inputPort := mu.NewMockUserInputPort(ctrl)
	mockCtrl := NewUserController(inputPort)

	const (
		callValidate int = iota
		callInputPort
	)
	type args struct {
		ctx *entity.Context
		in  *pb.DeleteUserRequest
	}
	tests := []struct {
		name       string
		args       args
		callExpect int
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			"correct",
			args{ct, &pb.DeleteUserRequest{
				UserId: "123412wq12q123q123q1",
			}},
			callInputPort,
			false,
		},
		{
			"validate user id error",
			args{ct, &pb.DeleteUserRequest{
				UserId: "123412wq12q123q123q1123",
			}},
			callValidate,
			true,
		},
		{
			"input port error",
			args{ct, &pb.DeleteUserRequest{
				UserId: "123412wq12q123q123q1",
			}},
			callInputPort,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputPort.EXPECT().DeleteUser(gomock.Any(), gomock.Any()).Do(func(ctx *entity.Context, in *dto.DeleteUserInputDto) {
				require.Equal(t, tt.args.in.UserId, in.UserId)
			}).Return(IfElse(tt.wantErr == true, fmt.Errorf("dummy error"), nil)).
				Times(IfElse(tt.callExpect == callValidate, 0, 1))

			if err := mockCtrl.DeleteUser(tt.args.ctx, tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("UserController.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
