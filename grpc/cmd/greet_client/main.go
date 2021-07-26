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
	//doServerStreaming(c)
	//doClientStreaming(c)
	doBiDiStreaming(c)	
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

func doClientStreaming(c greet.GreetServiceClient){
	fmt.Println("Starting to do a Client Streaming RPC ...")
	requests := []*greet.LongGreetRequest{
		&greet.LongGreetRequest{
			Greeting: &greet.Greeting{
				FirstName: "Allan",
			},
		},
		&greet.LongGreetRequest{
			Greeting: &greet.Greeting{
				FirstName: "Chepkoy",
			},
		},
		&greet.LongGreetRequest{
			Greeting: &greet.Greeting{
				FirstName: "Manami",
			},
		},
	}
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v", err)
	}

	// Iterate over the slice and send each message individually
	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while recieving response from LongGreet: %v", err)
	}
	fmt.Printf("LongGreet Response: %v\n", res)

}

func doBiDiStreaming(c greet.GreetServiceClient){
	fmt.Println("Starting to do a BIDI Streaming RPC ...")
	//create a stream by invoking the client
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
		return
	}

	requests := []*greet.GreetEveryoneRequest{
		&greet.GreetEveryoneRequest{
			Greeting: &greet.Greeting{
				FirstName: "Allan",
			},
		},
		&greet.GreetEveryoneRequest{
			Greeting: &greet.Greeting{
				FirstName: "Chepkoy",
			},
		},
		&greet.GreetEveryoneRequest{
			Greeting: &greet.Greeting{
				FirstName: "Manami",
			},
		},
	}

	waitc := make(chan struct{})
	// send a bunch of messages to the client (go routine)
	go func ()  {
		// function to send a bunch of messages
		for _, req := range requests {
			fmt.Printf("Sending Message: %v\n", req)
			stream.Send(req)
		}
		stream.CloseSend()
		
	}()
	//we recieve a bunch of messages from the client (go routine)
	go func ()  {
		// function to recieve a bunch of messages
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while recieving: %v", err)
				break
			}
			fmt.Printf("Recieved: %v\n", res.GetResult())
		}
		close(waitc)
		}()
	// block untill everything is done
	<-waitc
	
}