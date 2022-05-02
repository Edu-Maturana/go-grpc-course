package main

import (
	"context"
	"log"

	pb "github.com/Edu-Maturana/grpc-go-course/greet/proto"
)

func doGreet(client pb.GreetServiceClient) {
	log.Print("doGreet function was invoked")
	res, err := client.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Edu",
	})

	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}

	log.Printf("Greeting: %s", res.Result)
}
