package main

import (
	"io"
	"log"

	pb "github.com/f2hard3/grpc-go/calculator/proto"
)

func (*Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Printf("Max function was invoked")

	var maximum int32 = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream %v\n", err)
		}

		if number := req.Number; number > maximum {
			maximum = number

			log.Printf("Receiving number: %d\n", number)

			err := stream.Send(&pb.MaxResponse{
				Result: maximum,
			})

			if err != nil {
				log.Fatalf("Error while sending response: %v\n", err)
			}
		}

	}
}
