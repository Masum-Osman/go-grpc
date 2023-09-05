package main

import (
	"context"
	"log"

	pb "go-grpc/greet/proto"
)

// func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest, opts ...grpc.CallOption) (*pb.GreetResponse, error) {
// 	log.Printf("Greet function was invoked with %v\n", in)

// 	return &pb.GreetResponse{
// 		Result: "Hello " + in.FirstName,
// 	}, nil
// }

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v\n", in)

	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}

// Greet(context.Context, *GreetRequest) (*GreetResponse, error)
