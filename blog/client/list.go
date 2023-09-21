package main

import (
	"context"
	pb "go-grpc/blog/proto"
	"io"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlog(c pb.BlogServiceClient) {
	log.Println("List blog was invoked ------------ ")

	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Failed to List blog : %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something Happened: %v\n", err)
		}

		log.Println(res)
	}
}
