package main

import (
	"log"

	pb "go-grpc/calculator/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to listen on : %v\n", err)
	}

	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)

	// doSum(c)
	// doPrimes(c)
	// doAvg(c)
	doMax(c)
}
