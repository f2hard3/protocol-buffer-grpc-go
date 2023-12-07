package main

import (
	"context"
	"fmt"

	pb "github.com/f2hard3/grpc-go/calculator/proto"
)

func (*Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	fmt.Printf("Sum function was invoked with %v\n", in)

	return &pb.SumResponse{
		Result: in.FirstNumber + in.SecondNumber,
	}, nil
}
