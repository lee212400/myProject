package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	pb "github.com/lee212400/myProject/rpc"
)

type server struct {
	pb.UnimplementedSampleServiceServer
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	fmt.Println("GetUser API Start ---------")
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "No found metadata")
	}

	auth := md["authorization"]
	lang := md["lang"]
	reqId := md["x-request-id"]
	traceId := md["x-trace-id"]

	if len(auth) == 0 {
		return nil, status.Error(codes.Unauthenticated, "No authorization")
	}

	if len(lang) == 0 {
		return nil, status.Error(codes.Unauthenticated, "No authorization")
	}

	tk := strings.Split(auth[0], "Bearer ")

	if len(tk) < 2 {
		return nil, status.Error(codes.Unauthenticated, "Invaild token")
	}

	fmt.Println("requst header token:", tk[1])
	fmt.Println("requst header lang:", lang[0])
	fmt.Println("requst header reqId:", reqId[0])
	fmt.Println("requst header traceId:", traceId[0])

	return &pb.GetUserResponse{
		User: &pb.User{
			Email:     "test@test.com",
			FirstName: "fist_name",
			LastName:  "last_name",
		},
	}, nil
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer(
		grpc.UnaryInterceptor(unaryInterceptor),
	)
	pb.RegisterSampleServiceServer(s, &server{})
	log.Println("Server running at :50051")
	s.Serve(lis)
}

func getData(ctx *context.Context) (map[string]string, error) {
	// db,外部API処理
	res := map[string]string{
		"eame":  "name",
		"email": "sample@test.com",
	}
	return res, nil
}

func unaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	fmt.Println("Interceptor start ---------")
	// API service FullMethod
	fullMethod := info.FullMethod // "/ServiceName/MethodName"
	path := strings.Split(fullMethod, "/")

	// 実行されるメソッド名
	fmt.Println("API Method Called:", path[len(path)-1])

	// クライアント情報抽出
	p, ok := peer.FromContext(ctx)
	if ok {
		fmt.Println("Client IP:", p.Addr)
	}

	// requestデータ
	fmt.Println("request::::", req)

	// grpcメソッド呼び出す
	return handler(ctx, req)
}
