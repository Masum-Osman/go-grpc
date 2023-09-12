package main

import (
	pb "go-grpc/greet/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on : %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	opts := []grpc.ServerOption{}
	tls := true

	if tls {
		certFile := "ssl/server.crt"
		key := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, key)

		if err != nil {
			log.Fatalf("Failed to load certificate : %v\n", err)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)
	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(listen); err != nil {
		log.Fatalf("Failed to Serve: %v\n", err)
	}
}
