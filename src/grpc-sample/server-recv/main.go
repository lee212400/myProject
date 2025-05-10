package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/lee212400/myProject/rpc"
)

type server struct {
	pb.UnimplementedSampleServiceServer
}

func (s *server) ChatStream(stream pb.SampleService_ChatStreamServer) error {
	// clientのメッセージをすべて受信
	var allMessages string
	for {
		req, err := stream.Recv()
		if err == io.EOF { // 受信が終わったループ終了
			break
		}
		if err != nil {
			return fmt.Errorf("failed to receive a message: %v", err)
		}
		// responseデータ作成
		allMessages += req.GetMessage() + "\n"
	}

	// clinetにレスポンス送信
	response := &pb.ChatResponse{
		ResUser:    "Server",
		ResMessage: "Received the following messages:\n" + allMessages,
	}

	if err := stream.SendAndClose(response); err != nil {
		return fmt.Errorf("failed to send response: %v", err)
	}

	return nil
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()
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
