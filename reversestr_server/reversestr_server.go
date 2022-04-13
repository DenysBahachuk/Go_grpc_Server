package main

import (
	"context"
	"log"
	"net"

	pb "github.com/DenysBahachuk/go-reversestr-grpc/reversestr"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	pb.UnimplementedReverserServer
}

func reverseString(str string) string {
	result := ""
	for _, s := range str {
		result = string(s) + result
	}
	return result
}

func (s *GrpcServer) ReverseString(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Printf("Received: %v", req.GetStr())
	reversedString := reverseString(req.GetStr())
	return &pb.Response{Str: reversedString}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":4545")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterReverserServer(s, &GrpcServer{})
	log.Printf("Server is listening at %v", lis.Addr())

	err1 := s.Serve(lis)
	if err1 != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
