package main

import (
	"context"
	"fmt"
	"net"
	"log"
	"google.golang.org/grpc"
	"github.com/Masum-Osman/grpc-go/greetpb"
)

type server struct {}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreeResponse, error) {

	fmt.Printf("Greet function is invoked with %v", req)

	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &greetpb.GreeResponse { 
		Result : result,
	}
	
	return res, nil
}

func main()  {
	fmt.Println("Hello World")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Fatal to serve: %v", err)
	}
}	