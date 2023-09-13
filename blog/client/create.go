package main

import (
	"context"
	pb "go-grpc/blog/proto"
	"log"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("----- Create blog was invoked -----")

	blog := pb.Blog{
		AuthorId: "Masum",
		Title:    "Full Massage",
		Content:  "You know what I mean",
	}

	res, err := c.CreateBlog(context.Background(), &blog)
	if err != nil {
		log.Fatalf("Failed to post blog : %v\n", err)
	}

	log.Printf("Blog has been created%v\n", res.Id)
	return res.Id
}
