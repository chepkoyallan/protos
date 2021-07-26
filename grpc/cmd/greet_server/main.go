package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"

	// "time"

	"google.golang.org/grpc"

	"github.com/chepkoyallan/protos/pkg/greet"
	// "google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greet.GreetRequest) (*greet.GreetResponse, error){
	fmt.Printf("Greet function invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &greet.GreetResponse{
		Result: result,
	}
	return res, nil

}

func (*server) GreetManyTimes(req *greet.GreetManyTimesRequest, stream greet.GreetService_GreetManyTimesServer) error {
	fmt.Printf("GreetManyTime function invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstName + " number" + strconv.Itoa(i)
		res := &greet.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		//time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func (*server) LongGreet(stream greet.GreetService_LongGreetServer) error {
	fmt.Printf("Long greet function invoked with a streaming reqyest")
	result := "Hello "

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// finished reading client stream
			return stream.SendAndClose(&greet.LongGreetResponse{
				Result: result,
			})

		}
		if err != nil {
			log.Fatalf("error while reading client stream %v\n", err)
		}

		firstName := req.GetGreeting().GetFirstName()
		result += firstName + "! "
	}
}

func (*server) GreetEveryone(stream greet.GreetService_GreetEveryoneServer) error {
	fmt.Printf("Greet everyone function invoked with a streaming request\n")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err !=nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}
		firstName := req.GetGreeting().GetFirstName()
		result := "Hello " + firstName + "! "

		sendErr := stream.Send(&greet.GreetEveryoneResponse{
			Result: result,
		})

		if sendErr != nil {
			log.Fatalf("Error while sending data to client %v", err)
		}
	}
}

func main() {
	port := 50051
	fmt.Printf("GRPC Server started at port: %v\n", port)

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