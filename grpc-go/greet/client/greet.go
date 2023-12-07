package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/f2hard3/grpc-go/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Sunggon",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	fmt.Printf("Greeting %s\n", res.Result)
}
