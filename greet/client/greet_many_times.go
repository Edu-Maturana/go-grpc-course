package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Edu-Maturana/grpc-go-course/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Print("doGreetManyTimes function was invoked")

	req := pb.GreetRequest{
		FirstName: "Edu",
	}

	stream, err := c.GreetManyTimes(context.Background(), &req)

	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes RPC: %v", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}

		log.Printf("Response from GreetManyTimes: %s", msg.Result)
	}
}
