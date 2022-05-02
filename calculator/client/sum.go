package main

import (
	"context"
	"log"

	pb "github.com/Edu-Maturana/grpc-go-course/calculator/proto"
)

func sum(client pb.SumServiceClient) {
	log.Print("doGreet function was invoked")
	res, err := client.Sum(context.Background(), &pb.CalculatorRequest{
		Arguments: []int64{3, 10},
	})

	if err != nil {
		log.Fatalf("Error while calling sum RPC: %v", err)
	}

	log.Printf("Result: %v", res.Value)
}
