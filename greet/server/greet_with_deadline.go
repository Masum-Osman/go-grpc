package main

import (
	"context"
	pb "go-grpc/greet/proto"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Println("GreetWithDeadline was invoked")

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("The client cancelled the request!")
			return nil, status.Error(codes.Canceled, "The client canceled the request")
		}

		time.Sleep(1 * time.Second)
	}

	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
