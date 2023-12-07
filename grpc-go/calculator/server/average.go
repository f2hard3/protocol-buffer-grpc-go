package main

import (
	"io"
	"log"

	pb "github.com/f2hard3/grpc-go/calculator/proto"
)

func (*Server) Average(stream pb.CalculatorService_AverageServer) error {
	log.Printf("Average function was invoked")

	var sum, cnt int32 = 0, 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AverageResponse{
				Result: float64(sum) / float64(cnt),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream %v\n", err)
		}

		log.Printf("Receiving number: %d\n", req.Number)
		sum += req.Number
		cnt++
	}
}
