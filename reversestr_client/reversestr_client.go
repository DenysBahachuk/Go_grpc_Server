package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter a string (or press 'q' to exit): ")
		str, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error reading input: %v", err)
			os.Exit(2)
		}
		if str == "q\r\n" {
			fmt.Print("Goodbye!")
			break
		}

		response, err1 := client.ReverseString(ctx, &pb.Request{Str: str})
		if err1 != nil {
			log.Printf("Failed to create a user: %v", err1)
		}
		log.Printf("Reversed string: %v", response.Str)
	}
}
