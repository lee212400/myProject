package main

import (
	"context"
	"log"
	"net"

	"github.com/lee212400/myProject/interface/controller"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/lee212400/myProject/register"
	pb "github.com/lee212400/myProject/rpc/proto"
	uc "github.com/lee212400/myProject/utils/context"
)

type userService struct {
	userController *controller.UserController
	pb.UnimplementedUserServiceServer
}

func (s *userService) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	newCtx := uc.NewContext(ctx)
	err := s.userController.GetUser(newCtx, in)
	if err != nil {
		return &pb.GetUserResponse{}, err
	}
	return newCtx.Response.(*pb.GetUserResponse), nil
}
func (s *userService) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	newCtx := uc.NewContext(ctx)
	err := s.userController.CreateUser(newCtx, in)
	if err != nil {
		return &pb.CreateUserResponse{}, err
	}
	return newCtx.Response.(*pb.CreateUserResponse), nil
}
func (s *userService) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	newCtx := uc.NewContext(ctx)
	err := s.userController.UpdateUser(newCtx, in)
	if err != nil {
		return &pb.UpdateUserResponse{}, err
	}
	return newCtx.Response.(*pb.UpdateUserResponse), nil
}
func (s *userService) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	newCtx := uc.NewContext(ctx)
	err := s.userController.DeleteUser(newCtx, in)
	if err != nil {
		return &pb.DeleteUserResponse{}, err
	}
	return newCtx.Response.(*pb.DeleteUserResponse), nil
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()

	userService := &userService{
		userController: register.UserInit(),
	}

	pb.RegisterUserServiceServer(s, userService)
	log.Println("Server running at :50051")

	reflection.Register(s)

	s.Serve(lis)
}
