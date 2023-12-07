package main

import (
	"context"
	"io"
	"log"

	pb "github.com/f2hard3/grpc-go/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlogs(c pb.BlogServiceClient) {
	log.Println("listBlogs was invoked")

	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Failed to list blogs: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Failed to list blogs: %v\n", err)
		}

		log.Println(res)
	}

}
