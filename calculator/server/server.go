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

	lis, err := net.Listen("tcp", IP+":"+PORT)
	if nil != err {
		log.Fatalf("Error while create listen %v", err)
	}

	s := grpc.NewServer()

	calculatorpb.RegisterCalculatorServer(s, &Server{})

	fmt.Println("Calculator is running")
	err = s.Serve(lis)

	if nil != err {
		log.Fatalf("Error while serve %v", err)
	}
}
