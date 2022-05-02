package main

import (
	"context"
	"log"

	pb "github.com/Edu-Maturana/grpc-go-course/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Print("Sum function was invoked")
	return &pb.CalculatorResponse{
		Result: in.FirstValue + in.SecondValue,
	}, nil
}
