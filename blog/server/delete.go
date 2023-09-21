package main

import (
	"context"
	"fmt"
	pb "go-grpc/blog/proto"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteBlog(ctx context.Context, in *pb.BlogId) (*empty.Empty, error) {
	log.Print("Deleting blog was invoked with %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID: %v\n", err))

	}

	res, err := collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot delete from mongo DB: %v\n", err))

	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Not found: %v\n", err))
	}

	return &emptypb.Empty{}, nil
}
