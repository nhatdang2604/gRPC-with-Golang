package main

import (
	"context"
	"log"

	"github.com/nhatdang2604/gRPC-with-Golang/calculator/calculatorpb"
	"google.golang.org/grpc"
)

const (
	IP   = "localhost"
	PORT = "50030"
)

func callSum(client calculatorpb.CalculatorClient) {

	log.Println("Calling Sum API")

	response, err := client.Sum(context.Background(), &calculatorpb.SumRequest{
		Num1: 7,
		Num2: 6,
	})

	if nil != err {
		log.Fatalf("Call Sum API error: %v", err)
		return
	}

	log.Printf("Sum API responsed: %v", response.GetResult())
}

func main() {
	clientConnection, err := grpc.Dial(IP+":"+PORT, grpc.WithInsecure())

	//Error handle
	if nil != err {
		log.Fatalf("Error while dial %v", err)
	}

	//Closing the connection after using
	defer clientConnection.Close()

	client := calculatorpb.NewCalculatorClient(clientConnection)

	//log.Printf("Server client %f", client)

	callSum(client)
}
