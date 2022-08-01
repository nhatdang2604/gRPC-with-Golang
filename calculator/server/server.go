package main

import (
	"context"
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

func (server *Server) Sum(ctx context.Context, request *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {

	log.Println("Sum API called ...")
	response := &calculatorpb.SumResponse{
		Result: request.GetNum1() + request.GetNum2(),
	}

	return response, nil
}

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
