package main

import (
	"fmt"
	"log"

	"github.com/chepkoyallan/protos/protos"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello from greet client")
	// create connection to the server
	cc, err := grpc.Dial("localhost:500051", grpc.WithInsecure())
	if err != nil{
		log.Fatalf("coluld not connect %v", err)
	}

	defer cc.Close()

	c := protos.NewGreetServiceClient(cc)
	fmt.Printf("Created Client: %f", c)
}