package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/f2hard3/grpc-go/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("createBlog was invoked")

	blog := &pb.Blog{
		AuthorId: "Sunggon",
		Title:    "My first blog",
		Content:  "I love playing tennis",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Failed to create blog: %v\n", err)
	}

	fmt.Printf("Blog has been created: %s\n", res.Id)

	return res.Id
}
