package main

import (
	"context"
	pb "go-grpc/greet/proto"
	"log"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked.")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Masum",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Println("Greeting: %s\n", res)
}
