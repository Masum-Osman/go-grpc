package main

import ( 
	"context"
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
	// fmt.Println("New Client Connect: %f", c)

	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient)  {

	fmt.Println("Starting to do a Unary RPC...")

	req := &greetpb.GreetRequest {
		Greeting : &greetpb.Greeting {
			FirstName : "Osman",
			LastName : "Masum",
		},
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}
	log.Printf("Response From Greet: %v", res.Result)
}