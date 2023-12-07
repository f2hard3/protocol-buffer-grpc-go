package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/f2hard3/grpc-go/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  10,
		SecondNumber: 50,
	})

	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}

	fmt.Printf("Sum %d\n", res.Result)
}
