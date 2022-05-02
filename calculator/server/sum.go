package main

import (
	"context"
	"log"

	pb "github.com/Edu-Maturana/grpc-go-course/calculator/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Printf("Greet function was invoked with %v", in)
	return &pb.CalculatorResponse{
		Value: in.Arguments[0] + in.Arguments[1],
	}, nil
}
