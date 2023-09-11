package main

import (
	"context"
	pb "go-grpc/calculator/proto"
	"io"
	"log"
	"time"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doGreetEveryone was invoked")

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	waitc := make(chan struct{})

	go func() {
		numbers := []int32{4, 7, 2, 19, 4, 6, 32}

		for _, number := range numbers {
			log.Printf("Sending Number : %d\n", number)

			stream.Send(&pb.MaxRequest{
				Number: number,
			})
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Error while receiving stream: %v\n", err)
				break
			}

			log.Printf("Received: %v\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}
