package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Edu-Maturana/grpc-go-course/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Print("doLongGreet function was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Edu"},
		{FirstName: "Andres"},
		{FirstName: "Juan"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response: %v", err)
	}

	log.Printf("LongGreet response: %v", res)
}
