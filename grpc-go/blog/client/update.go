package main

import (
	"context"
	"log"

	pb "github.com/f2hard3/grpc-go/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("updateBlog was invoked")

	newblog := &pb.Blog{
		Id:       id,
		AuthorId: "Jigyo",
		Title:    "How old are you?",
		Content:  "I am 3 years old",
	}

	_, err := c.UpdateBlog(context.Background(), newblog)

	if err != nil {
		log.Fatalf("Failed to update blog: %v\n", err)
	}

	log.Println("Blog has been updated")
}
