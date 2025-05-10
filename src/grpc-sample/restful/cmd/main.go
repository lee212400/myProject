package main

import (
	"context"
	"fmt"
	"log"
	"net"

	protovalidate "github.com/bufbuild/protovalidate-go"
	pb "github.com/lee212400/myProject/rpc"
	"google.golang.org/grpc"
)

var myValidate protovalidate.Validator

func init() {
	v, _ := protovalidate.New()
	myValidate = v
}

type server struct {
	pb.UnimplementedSampleServiceServer
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	if err := myValidate.Validate(req); err != nil {
		return &pb.GetUserResponse{}, err
	}
	fmt.Println("GetUser User OK:")
	return &pb.GetUserResponse{
		User: &pb.User{
			Email:     "test@test.com",
			FirstName: "first",
			LastName:  "last",
		},
	}, nil
}

func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	if err := myValidate.Validate(req); err != nil {
		return &pb.CreateUserResponse{}, err
	}
	fmt.Println("Email::", req.Email)
	email := req.Email
	fm := req.FirstName
	lm := req.LastName

	fmt.Println("CreateUser User OK:")
	return &pb.CreateUserResponse{
		User: &pb.User{
			Email:     email,
			FirstName: fm,
			LastName:  lm,
		},
	}, nil
}

func (s *server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	nm := req.Name
	fmt.Println("UpdateUser User OK:", nm)
	return &pb.UpdateUserResponse{
		User: &pb.User{
			Email:     "test@test.com",
			FirstName: nm,
			LastName:  nm,
		},
	}, nil
}

func (s *server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	id := req.UserId

	fmt.Println("Delete User OK:", id)
	return &pb.DeleteUserResponse{}, nil
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer(grpc.UnaryInterceptor(RecoveryInterceptor))
	pb.RegisterSampleServiceServer(s, &server{})
	log.Println("Server running at :50051")
	s.Serve(lis)
}

func RecoveryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("[Recovery] panic recovered: %v\n", r)
			err = fmt.Errorf("internal server error") // 에러 반환
		}
	}()
	return handler(ctx, req)
}
