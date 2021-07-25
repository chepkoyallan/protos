package main

import (
	"context"
	"fmt"
	"github.com/chepkoyallan/protos/pkg/greet"
	"google.golang.org/grpc"
	"log"
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