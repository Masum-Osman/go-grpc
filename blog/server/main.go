package main

import (
	"context"
	pb "go-grpc/blog/proto"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"
var collection *mongo.Collection

type Server struct {
	pb.BlogServiceServer
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatalf("Failed to Connect MongoDB on : %v\n", err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatalf("Failed to Connect MongoDB on : %v\n", err)
	}

	collection = client.Database("blogdb").Collection("blog")

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on : %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &Server{})

	if err = s.Serve(listen); err != nil {
		log.Fatalf("Failed to Serve: %v\n", err)
	}
}
