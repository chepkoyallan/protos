package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"

	"github.com/chepkoyallan/protos/pkg/greet"
	// "google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greet.GreetRequest) (*greet.GreetResponse, error){
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &greet.GreetResponse{
		Result: result,
	}
	return res, nil

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
	greet.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}