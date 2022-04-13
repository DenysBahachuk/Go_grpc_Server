package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/DenysBahachuk/go-reversestr-grpc/reversestr"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4545", grpc.WithInsecure())
	if err != nil {
		log.Printf("Connection failed: %v", err)
		os.Exit(1)
	}
	defer conn.Close()
	client := pb.NewReverserClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	input := "Hello"

	response, err1 := client.ReverseString(ctx, &pb.Request{Str: input})

	if err1 != nil {
		log.Printf("Failed to create a user: %v", err1)
	}

	log.Printf("Reversed string: %v", response.Str)
}
