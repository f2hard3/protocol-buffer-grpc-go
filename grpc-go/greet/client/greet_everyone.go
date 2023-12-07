package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/f2hard3/grpc-go/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Printf("doGreetEveryone was invoked")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Sunggon"},
		{FirstName: "Jigyo"},
		{FirstName: "Youri"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sending req: %v\n", req)
			stream.Send(req)
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
				log.Fatalf("Error while receiving response from GreetEveryone: %v\n", err)
				break
			}

			log.Printf("Receiving res: %v\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}
