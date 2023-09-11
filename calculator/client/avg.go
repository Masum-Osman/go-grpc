package main

import (
	"context"
	pb "go-grpc/calculator/proto"
	"log"
	"time"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Println("doAvg was invoked")

	// reqs := []*pb.AvgRequest{
	// 	{Number: 1},
	// 	{Number: 2},
	// 	{Number: 3},
	// 	{Number: 4},
	// }

	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatalf("Error while calling Avg %v\t", err)
	}

	numbers := []int32{3, 5, 9, 54, 23}

	for _, number := range numbers {
		log.Printf("Sending number: %d\n", number)

		stream.Send(&pb.AvgRequest{
			Number: number,
		})
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from Avg: %v\n", err)
	}

	log.Printf("Avg: %f\n", res.Result)
}