package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/f2hard3/grpc-go/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("readBlog was invoked")

	req := &pb.BlogId{Id: id}
	res, err := c.ReadBlog(context.Background(), req)

	if err != nil {
		log.Fatalf("Failed to read blog: %v\n", err)
	}

	fmt.Printf("Blog was read: %s\n", res)

	return res
}
