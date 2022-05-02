package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/Edu-Maturana/grpc-go-course/greet/proto"
)

func (s *server) LongGreet(stream GreetService_LongGreetServer) error {
	log.Print("LongGreet function was invoked with success")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		res += fmt.Sprintf("Hello %s! ", req.FirstName)
	}
}
