package main

import (
	"fmt"
	"log"
	"google.golang.org/grpc"
	"github.com/Masum-Osman/grpc-go/greetpb"
)

func main()  {
	fmt.Println("Hello, I'm a client...")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc) 
	fmt.Println("New Client Connect: %f", c)
}