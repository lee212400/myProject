package controller

import (
	"testing"

	"github.com/lee212400/myProject/domain/entity"
	pb "github.com/lee212400/myProject/rpc/proto"
	"github.com/lee212400/myProject/usecase"
)

func TestUserController_GetUser(t *testing.T) {
	type fields struct {
		inputPort usecase.UserInputPort
	}
	type args struct {
		ctx *entity.Context
		in  *pb.GetUserRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &UserController{
				inputPort: tt.fields.inputPort,
			}
			if err := i.GetUser(tt.args.ctx, tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("UserController.GetUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
