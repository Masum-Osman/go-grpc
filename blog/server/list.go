package main

import (
	"context"
	"fmt"
	pb "go-grpc/blog/proto"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListBlogs(in *empty.Empty, stream pb.BlogService_ListBlogsServer) error {
	log.Println("List Blog was invoked")

	cur, err := collection.Find(context.Background(), primitive.D{{}})
	if err != nil {
		return status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID: %v\n", err))

	}

	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		data := &BlogItem{}
		err := cur.Decode(data)

		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while decoding data from MongoDB: %", err),
			)
		}

		stream.Send(documentToBlog(data))
	}

	if err = cur.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %", err),
		)
	}

	return nil
}
