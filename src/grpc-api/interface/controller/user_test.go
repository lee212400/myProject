package controller

import (
	"context"
	"testing"

	"github.com/lee212400/myProject/domain/entity"
	mu "github.com/lee212400/myProject/mock/usecase"
	pb "github.com/lee212400/myProject/rpc/proto"
	"github.com/lee212400/myProject/usecase/dto"
	uc "github.com/lee212400/myProject/utils/context"
	"go.uber.org/mock/gomock"
)

func TestUserController_GetUser(t *testing.T) {

	ct := uc.NewContext(context.Background())
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	inputPort := mu.NewMockUserInputPort(ctrl)
	mockCtrl := NewUserController(inputPort)

	type args struct {
		ctx *entity.Context
		in  *pb.GetUserRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"case1",
			args{ct, &pb.GetUserRequest{
				UserId: "123412wq12q123q123q1",
			}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputPort.EXPECT().GetUser(gomock.Any(), gomock.Any()).Do(func(ctx *entity.Context, in *dto.GetUserInputDto) {
				t.Log("in:::", in)
			}).Return(nil)

			if err := mockCtrl.GetUser(tt.args.ctx, tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("UserController.GetUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
