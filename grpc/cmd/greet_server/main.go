package main

import (
	"fmt"
	"log"
	"net"

	"github.com/chepkoyallan/protos/github.com/chepkoyallan/protos"
	"google.golang.org/grpc"
)

type server struct{
	protos.UnimplementedGreetServiceServer
}

func main() {
	fmt.Println("Hello from greet server")

	// 1. Write a listener
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil{
		log.Fatalf("Failed to listen:  %v", err)
	}

	// 2. We need to register a servive
	s :=  grpc.NewServer()
	protos.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}