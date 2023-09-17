package main

import (
	"context"
	pb "go-grpc/blog/proto"
	"log"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("--- readBlog was invoked ---")

	req := &pb.BlogId{
		Id: id,
	}

	res, err := c.ReadBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to read blog : %v\n", err)
	}

	log.Printf("Blog was read: %v\n", res)

	return res
}
