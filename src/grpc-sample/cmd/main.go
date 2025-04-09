package main

import (
	"context"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"

	pb "github.com/lee212400/myProject/rpc"
)

type server struct {
	pb.UnimplementedSampleServiceServer
}

func (s *server) GetData(req *pb.StreamRequest, stream pb.SampleService_GetDataServer) error {
	ctx := stream.Context()
	dataCh := make(chan map[string]string, 1)
	errCh := make(chan error, 1)

	// GetData関連実行
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		data, err := getData(&ctx)
		if err != nil {
			errCh <- err
			return
		}
		dataCh <- data

	}()

	go func() {
		wg.Wait()
		close(dataCh)
	}()

	for {
		select {
		case <-ctx.Done():
			log.Println("context canceled or timeout:", ctx.Err())
			return ctx.Err()

		case err := <-errCh:
			log.Println("getData failed:", err)
			return err

		case data, ok := <-dataCh:
			if !ok {
				log.Println("Data channel closed, ending stream")
				return nil
			}

			res := &pb.StreamResponse{
				Name:  data["name"],
				Email: data["email"],
			}
			if err := stream.Send(res); err != nil {
				log.Printf("stream send error: %v", err)
				return err
			}
		}
	}

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
