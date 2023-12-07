package main

import (
	"context"
	"log"

	pb "github.com/f2hard3/grpc-go/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("deleteBlog was invoked")

	req := &pb.BlogId{Id: id}
	_, err := c.DeleteBlog(context.Background(), req)

	if err != nil {
		log.Fatalf("Failed to delete blog: %v\n", err)
	}

	log.Printf("Blog deleted with id: %v\n", id)
}
