package main

import (
	"context"
	pb "go-grpc/blog/proto"
	"log"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("Delete blog was invoked ---------")

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})
	if err != nil {
		log.Fatalf("Error from client on deleting blog %v\n", err)
	}

	log.Println("Blog was deleted")
}
