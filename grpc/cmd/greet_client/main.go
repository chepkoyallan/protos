package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/chepkoyallan/protos/pkg/greet"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello from greet client")
	// create connection to the server
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil{
		log.Fatalf("coluld not connect %v", err)
	}

	defer cc.Close()



	c := greet.NewGreetServiceClient(cc)

	// unary
	// doUnary(c)

	//server Streaming
	doServerStreaming(c)
	
}

func doUnary(c greet.GreetServiceClient){
	fmt.Println("Starting to do a  Unary RPC ...")
	req := &greet.GreetRequest{
		Greeting: &greet.Greeting{
			FirstName: "Allan",
			SecondName: "Chepkoy",
		},
	}
	res, err := c.Greet(context.Background(),req)
	if err != nil {
		log.Fatalf("Error while calling the greet RPC: %v", err)
	}
	log.Printf("Responsee from greeting: %v", res.Result)
}

func doServerStreaming(c greet.GreetServiceClient){
	fmt.Println("Starting to do a  Server Streaming RPC ...")
	req := &greet.GreetManyTimesRequest{
		Greeting: &greet.Greeting{
			FirstName: "Allan",
			SecondName: "Chepkoy",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling the Greet Many Times RPC: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// Reached end of stream 
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}
		log.Printf("Responsee from greeting: %v", msg.GetResult())
	}
}