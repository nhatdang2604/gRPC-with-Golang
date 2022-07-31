package main

import (
	"log"

	"github.com/nhatdang2604/gRPC-with-Golang/calculator/calculatorpb"
	"google.golang.org/grpc"
)

const (
	IP   = "localhost"
	PORT = "50030"
)

func main() {
	clientConnection, err := grpc.Dial(IP+":"+PORT, grpc.WithInsecure())

	//Error handle
	if nil != err {
		log.Fatalf("Error while dial %v", err)
	}

	//Closing the connection after using
	defer clientConnection.Close()

	client := calculatorpb.NewCalculatorClient(clientConnection)

	log.Printf("Server client %f", client)
}
