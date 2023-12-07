package main

import (
	"context"
	"log"

	pb "github.com/f2hard3/grpc-go/calculator/proto"
)

func doAverage(c pb.CalculatorServiceClient) {
	log.Printf("doAverage was invoked")

	stream, err := c.Average(context.Background())

	if err != nil {
		log.Fatalf("Error while opening the stream: %v\n", err)
	}

	numbers := []int32{13, 231, 568, 21, 88}

	for _, number := range numbers {
		log.Printf("Sending number: %d\n", number)

		stream.Send(&pb.AverageRequest{
			Number: number,
		})
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from Average: %v\n", err)
	}

	log.Printf("Average: %f\n", res.Result)
}
