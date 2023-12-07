package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/f2hard3/grpc-go/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Printf("doMax was invoked")

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	waitc := make(chan struct{})

	go func() {
		numbers := []int32{4, 7, 3, 12, 2, 6, 42}

		for _, number := range numbers {
			log.Printf("Sending number: %v\n", number)
			stream.Send(&pb.MaxRequest{
				Number: number,
			})
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error while reading stream: %v\n", err)
				break
			}

			log.Printf("Received a new maximum: %d\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}
