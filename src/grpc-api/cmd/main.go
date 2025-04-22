package main

import (
	"context"
	"log"
	"net"

	"github.com/lee212400/myProject/domain/entity"
	"github.com/lee212400/myProject/interface/controller"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/lee212400/myProject/register"
	pb "github.com/lee212400/myProject/rpc/proto"
	uc "github.com/lee212400/myProject/utils/context"
	env "github.com/lee212400/myProject/utils/env"
	ue "github.com/lee212400/myProject/utils/errors"
)

type userService struct {
	userController *controller.UserController
	pb.UnimplementedUserServiceServer
}

func (s *userService) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	newCtx := uc.NewContext(ctx)
	return handler[*pb.GetUserRequest, *pb.GetUserResponse](newCtx, in, s.userController.GetUser)
}

func (s *userService) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	newCtx := uc.NewContext(ctx)
	return handler[*pb.CreateUserRequest, *pb.CreateUserResponse](newCtx, in, s.userController.CreateUser)
}

func (s *userService) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	newCtx := uc.NewContext(ctx)
	return handler[*pb.UpdateUserRequest, *pb.UpdateUserResponse](newCtx, in, s.userController.UpdateUser)
}

func (s *userService) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	newCtx := uc.NewContext(ctx)
	return handler[*pb.DeleteUserRequest, *pb.DeleteUserResponse](newCtx, in, s.userController.DeleteUser)
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()

	userService := &userService{
		userController: register.UserInit(),
	}

	host := env.GetEnvString("MYSQL_HOST")
	user := env.GetEnvString("MYSQL_USER")
	pw := env.GetEnvString("MYSQL_PASSWORD")
	po := env.GetEnvString("MYSQL_PORT")
	db := env.GetEnvString("MYSQL_DATABASES")
	log.Printf("host:%s, user:%s, password:%s, port:%s, database:%s", host, user, pw, po, db)

	pb.RegisterUserServiceServer(s, userService)
	log.Println("Server running at :50051")

	reflection.Register(s)

	s.Serve(lis)
}

func handler[T, R any](ctx *entity.Context, in T, f func(ctx *entity.Context, in T) error) (R, error) {
	var res R
	err := f(ctx, in)
	if err, ok := err.(*ue.AppError); ok && err != nil {
		return res, err.Generate(ctx)
	}

	return ctx.Response.(R), nil
}
