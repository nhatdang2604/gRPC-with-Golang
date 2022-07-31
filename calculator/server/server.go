package main

import (
	"fmt"
	"log"
	"net"

	"github.com/nhatdang2604/gRPC-with-Golang/calculator/calculatorpb"
	"google.golang.org/grpc"
)

const (
	IP   = "0.0.0.0"
	PORT = "50030"
)

type Server struct{}

func main() {

	listener, err := net.Listen("tcp", IP+":"+PORT)
	if nil != err {
		log.Fatalf("Error while create listen %v", err)
	}

	server := grpc.NewServer()

	calculatorpb.RegisterCalculatorServer(server, &Server{})

	fmt.Println("Calculator is running")
	err = server.Serve(listener)

	if nil != err {
		log.Fatalf("Error while serve %v", err)
	}
}
