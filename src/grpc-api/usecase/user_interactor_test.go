package usecase

import (
	"context"
	"fmt"
	"testing"

	"github.com/lee212400/myProject/domain/entity"
	mock "github.com/lee212400/myProject/mock/repository"
	mu "github.com/lee212400/myProject/mock/usecase"
	"github.com/lee212400/myProject/usecase/dto"
	"github.com/lee212400/myProject/utils"
	uc "github.com/lee212400/myProject/utils/context"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func testGetUserData() *entity.User {
	return &entity.User{
		UserId:    "1q2w3e4r5t6y7u8i1q2w",
		Email:     "test@test.com",
		FirstName: "first",
		LastName:  "last",
		Age:       30,
	}
}

func TestUserInteractor_GetUser(t *testing.T) {
	ct := uc.NewContext(context.Background())

	const (
		callGetUser int = iota
		callOutputPort
	)

	type args struct {
		ctx         *entity.Context
		in          *dto.GetUserInputDto
		getUserData *entity.User
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
			args{ct, &dto.GetUserInputDto{
				UserId: "1q2w3e4r5t6y7u8i1q2w",
			}, testGetUserData()},

			callOutputPort,
			false,
		},
		{
			"Get User Error",
			args{ct, &dto.GetUserInputDto{
				UserId: "1q2w3e4r5t6y7u8i1q2w",
			}, &entity.User{}},
			callGetUser,
			true,
		},
		{
			"OutputPort Error",
			args{ct, &dto.GetUserInputDto{
				UserId: "1q2w3e4r5t6y7u8i1q2w",
			}, testGetUserData()},
			callOutputPort,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			postRepo := mock.NewMockPostProcessRepository(ctrl)
			userRepo := mock.NewMockUserRepository(ctrl)
			ouputPort := mu.NewMockUserOutputPort(ctrl)

			interactor := NewUserInteractor(postRepo, userRepo, ouputPort)

			defer ctrl.Finish()

			var ppErr error

			postRepo.EXPECT().PostProcess(gomock.Any(), gomock.Any()).Do(func(ctx *entity.Context, err *error) {
				ppErr = *err
			})

			userRepo.EXPECT().GetUser(gomock.Any(), gomock.Any()).Do(func(ctx *entity.Context, userId string) {
				require.Equal(t, tt.args.in.UserId, userId)
			}).Return(tt.args.getUserData, utils.IfElse(tt.wantErr && tt.callExpect == callGetUser, fmt.Errorf("dummy error"), nil)).
				Times(utils.IfElse(tt.callExpect >= callGetUser, 1, 0))

			ouputPort.EXPECT().GetUser(gomock.Any(), gomock.Any()).Do(func(ctx *entity.Context, out *dto.GetUserOutputDto) {
				require.Equal(t, &dto.GetUserOutputDto{
					User: &entity.User{
						UserId:    tt.args.getUserData.UserId,
						Email:     tt.args.getUserData.Email,
						FirstName: tt.args.getUserData.FirstName,
						LastName:  tt.args.getUserData.LastName,
						Age:       tt.args.getUserData.Age,
					},
				}, out)
			}).Return(utils.IfElse(tt.wantErr && tt.callExpect == callOutputPort, fmt.Errorf("dummy error"), nil)).
				Times(utils.IfElse(tt.callExpect >= callOutputPort, 1, 0))

			err := interactor.GetUser(tt.args.ctx, tt.args.in)

			if (err != nil) != tt.wantErr {
				t.Errorf("UserInteractor.GetUser() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr {
				require.Equal(t, ppErr, err)
			}
		})
	}
}

func TestUserInteractor_CreateUser(t *testing.T) {
	ct := uc.NewContext(context.Background())

	const (
		callCreateUser int = iota
		callGetUser
		callOutputPort
	)

	inDt := &entity.User{
		Email:     "test@test.com",
		FirstName: "first",
		LastName:  "last",
		Age:       30,
	}

	type args struct {
		ctx           *entity.Context
		in            *dto.CreateUserInputDto
		retCreateUser string
		retGetUser    *entity.User
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
			args{ct, &dto.CreateUserInputDto{
				User: inDt,
			}, "1q2w3e4r5t6y7u8i9o1q", testGetUserData()},
			callOutputPort,
			false,
		},
		{
			"fail: Create User",
			args{ct, &dto.CreateUserInputDto{
				User: inDt,
			}, "", &entity.User{}},
			callCreateUser,
			true,
		},
		{
			"fail: Get User",
			args{ct, &dto.CreateUserInputDto{
				User: inDt,
			}, "1q2w3e4r5t6y7u8i9o1q", &entity.User{}},
			callGetUser,
			true,
		},
		{
			"fail: OutputPort",
			args{ct, &dto.CreateUserInputDto{
				User: inDt,
			}, "1q2w3e4r5t6y7u8i9o1q", testGetUserData()},
			callOutputPort,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			postRepo := mock.NewMockPostProcessRepository(ctrl)
			userRepo := mock.NewMockUserRepository(ctrl)
			ouputPort := mu.NewMockUserOutputPort(ctrl)

			interactor := NewUserInteractor(postRepo, userRepo, ouputPort)

			defer ctrl.Finish()

			var ppErr error

			postRepo.EXPECT().PostProcess(gomock.Any(), gomock.Any()).Do(func(ctx *entity.Context, err *error) {
				ppErr = *err
			})

			userRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Do(func(ctx *entity.Context, firstName string, lastName string, email string, age int32) {
				require.Equal(t, tt.args.in.User.FirstName, firstName)
				require.Equal(t, tt.args.in.User.LastName, lastName)
				require.Equal(t, tt.args.in.User.Email, email)
				require.Equal(t, tt.args.in.User.Age, age)
			}).Return(tt.args.retCreateUser, utils.IfElse(tt.wantErr && tt.callExpect == callCreateUser, fmt.Errorf("dummy error"), nil)).
				Times(utils.IfElse(tt.callExpect >= callCreateUser, 1, 0))

			userRepo.EXPECT().GetUser(gomock.Any(), gomock.Any()).Do(func(ctx *entity.Context, userId string) {
				require.Equal(t, tt.args.retCreateUser, userId)
			}).Return(tt.args.retGetUser, utils.IfElse(tt.wantErr && tt.callExpect == callGetUser, fmt.Errorf("dummy error"), nil)).
				Times(utils.IfElse(tt.callExpect >= callGetUser, 1, 0))

			ouputPort.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Do(func(ctx *entity.Context, out *dto.CreateUserOutputDto) {
				require.Equal(t, &dto.CreateUserOutputDto{
					User: &entity.User{
						UserId:    tt.args.retGetUser.UserId,
						Email:     tt.args.retGetUser.Email,
						FirstName: tt.args.retGetUser.FirstName,
						LastName:  tt.args.retGetUser.LastName,
						Age:       tt.args.retGetUser.Age,
					},
				}, out)
			}).Return(utils.IfElse(tt.wantErr && tt.callExpect == callOutputPort, fmt.Errorf("dummy error"), nil)).
				Times(utils.IfElse(tt.callExpect >= callOutputPort, 1, 0))

			err := interactor.CreateUser(tt.args.ctx, tt.args.in)

			if (err != nil) != tt.wantErr {
				t.Errorf("UserInteractor.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr {
				require.Equal(t, ppErr, err)
			}
		})
	}
}

func TestUserInteractor_UpdateUser(t *testing.T) {
	ct := uc.NewContext(context.Background())

	const (
		callUpdateUser int = iota
		callGetUser
		callOutputPort
	)

	type args struct {
		ctx        *entity.Context
		in         *dto.UpdateUserInputDto
		retGetUser *entity.User
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
			args{ct, &dto.UpdateUserInputDto{UserId: "1q2w3e4r5t6y7u8i9o1q", Age: 30},
				testGetUserData()},
			callOutputPort,
			false,
		},
		{
			"fail: update user",
			args{ct, &dto.UpdateUserInputDto{UserId: "1q2w3e4r5t6y7u8i9o1q", Age: 30},
				&entity.User{}},
			callUpdateUser,
			true,
		},
		{
			"fail: get user",
			args{ct, &dto.UpdateUserInputDto{UserId: "1q2w3e4r5t6y7u8i9o1q", Age: 30},
				&entity.User{}},
			callGetUser,
			true,
		},
		{
			"fail: output port",
			args{ct, &dto.UpdateUserInputDto{UserId: "1q2w3e4r5t6y7u8i9o1q", Age: 30},
				testGetUserData()},
			callOutputPort,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			postRepo := mock.NewMockPostProcessRepository(ctrl)
			userRepo := mock.NewMockUserRepository(ctrl)
			ouputPort := mu.NewMockUserOutputPort(ctrl)

			interactor := NewUserInteractor(postRepo, userRepo, ouputPort)

			defer ctrl.Finish()

			var ppErr error

			postRepo.EXPECT().PostProcess(gomock.Any(), gomock.Any()).Do(func(ctx *entity.Context, err *error) {
				ppErr = *err
			})

			userRepo.EXPECT().UpdateUser(gomock.Any(), gomock.Any(), gomock.Any()).Do(func(ctx *entity.Context, userId string, age int32) {
				require.Equal(t, tt.args.in.UserId, userId)
				require.Equal(t, tt.args.in.Age, age)
			}).Return(utils.IfElse(tt.wantErr && tt.callExpect == callUpdateUser, fmt.Errorf("dummy error"), nil)).
				Times(utils.IfElse(tt.callExpect >= callUpdateUser, 1, 0))

			userRepo.EXPECT().GetUser(gomock.Any(), gomock.Any()).Do(func(ctx *entity.Context, userId string) {
				require.Equal(t, tt.args.in.UserId, userId)
			}).Return(tt.args.retGetUser, utils.IfElse(tt.wantErr && tt.callExpect == callGetUser, fmt.Errorf("dummy error"), nil)).
				Times(utils.IfElse(tt.callExpect >= callGetUser, 1, 0))

			ouputPort.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Do(func(ctx *entity.Context, out *dto.UpdateUserOutputDto) {
				require.Equal(t, &dto.UpdateUserOutputDto{
					User: &entity.User{
						UserId:    tt.args.retGetUser.UserId,
						Email:     tt.args.retGetUser.Email,
						FirstName: tt.args.retGetUser.FirstName,
						LastName:  tt.args.retGetUser.LastName,
						Age:       tt.args.retGetUser.Age,
					},
				}, out)
			}).Return(utils.IfElse(tt.wantErr && tt.callExpect == callOutputPort, fmt.Errorf("dummy error"), nil)).
				Times(utils.IfElse(tt.callExpect >= callOutputPort, 1, 0))

			err := interactor.UpdateUser(tt.args.ctx, tt.args.in)

			if (err != nil) != tt.wantErr {
				t.Errorf("UserInteractor.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr {
				require.Equal(t, ppErr, err)
			}
		})
	}
}

func TestUserInteractor_DeleteUser(t *testing.T) {
	ct := uc.NewContext(context.Background())

	const (
		callDeleteUser int = iota
		callOutputPort
	)

	type args struct {
		ctx *entity.Context
		in  *dto.DeleteUserInputDto
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
			args{ct, &dto.DeleteUserInputDto{
				UserId: "1q2w3e4r5t6y7u8i9o1q",
			}}, callOutputPort, false,
		},
		{
			"fail: delete user",
			args{ct, &dto.DeleteUserInputDto{
				UserId: "1q2w3e4r5t6y7u8i9o1q",
			}}, callDeleteUser, true,
		},
		{
			"fail: output port",
			args{ct, &dto.DeleteUserInputDto{
				UserId: "1q2w3e4r5t6y7u8i9o1q",
			}}, callOutputPort, true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			postRepo := mock.NewMockPostProcessRepository(ctrl)
			userRepo := mock.NewMockUserRepository(ctrl)
			ouputPort := mu.NewMockUserOutputPort(ctrl)

			interactor := NewUserInteractor(postRepo, userRepo, ouputPort)

			defer ctrl.Finish()

			var ppErr error

			postRepo.EXPECT().PostProcess(gomock.Any(), gomock.Any()).Do(func(ctx *entity.Context, err *error) {
				ppErr = *err
			})

			userRepo.EXPECT().DeleteUser(gomock.Any(), gomock.Any()).Do(func(ctx *entity.Context, userId string) {
				require.Equal(t, tt.args.in.UserId, userId)
			}).Return(utils.IfElse(tt.wantErr && tt.callExpect == callDeleteUser, fmt.Errorf("dummy error"), nil)).
				Times(utils.IfElse(tt.callExpect >= callDeleteUser, 1, 0))

			ouputPort.EXPECT().DeleteUser(gomock.Any(), gomock.Any()).Do(func(ctx *entity.Context, out *dto.DeleteUserOutputDto) {
				require.Equal(t, &dto.DeleteUserOutputDto{}, out)
			}).Return(utils.IfElse(tt.wantErr && tt.callExpect == callOutputPort, fmt.Errorf("dummy error"), nil)).
				Times(utils.IfElse(tt.callExpect >= callOutputPort, 1, 0))

			err := interactor.DeleteUser(tt.args.ctx, tt.args.in)

			if (err != nil) != tt.wantErr {
				t.Errorf("UserInteractor.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr {
				require.Equal(t, ppErr, err)
			}
		})
	}
}
