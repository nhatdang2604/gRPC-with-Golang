package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"time"

	"github.com/nhatdang2604/gRPC-with-Golang/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
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

func (server *Server) PrimeNumberDecomposition(request *calculatorpb.PNDRequest,
	stream calculatorpb.Calculator_PrimeNumberDecompositionServer) error {

	//Logging
	log.Println("Prime Number Decomposition API called ...")

	//Get the number to decomposite
	number := request.GetNumber()

	//Prime number to decomposition
	prime := int32(2)

	//Algorithm to decomposition
	for number > 1 {
		if number%prime == 0 {
			number = number / prime

			//send the prime to the client
			stream.Send(&calculatorpb.PNDResponse{
				Number: prime,
			})

		} else {
			prime++
			log.Printf("Prime increase to %v", prime)
		}
	}

	return nil
}

func (server *Server) Average(stream calculatorpb.Calculator_AverageServer) error {

	//buffer to calculate the average
	var total float32
	var count int = 0

	//Logging
	log.Println("Average API called ...")

	for {
		request, err := stream.Recv()
		if io.EOF == err {

			//Calculate the average
			result := total / float32(count)
			response := &calculatorpb.AverageResponse{
				Result: result,
			}

			//Send the response and terminate the api
			return stream.SendAndClose(response)
		}

		if nil != err {
			log.Fatalf("Error while reciving average %v", err)
			return err
		}

		log.Printf("Recieved request: %v\n", request)
		count += 1
		total += request.GetNumber()
	}

}

func (server *Server) FindMax(stream calculatorpb.Calculator_FindMaxServer) error {

	//Logging
	log.Println("Find Max API called ...")

	max := int32(math.MinInt32)

	for {
		request, err := stream.Recv()

		//EOF handle
		if io.EOF == err {
			log.Printf("EOF....\n")
			return nil
		}

		//Error handle while recieving request
		if nil != err {
			log.Fatalf("Error while recieving find max: %v", err)
			return err
		}

		//Calculate the current max
		buffer := request.GetNumber()
		if max < buffer {
			max = buffer
		}

		//Send the current max to client
		err = stream.Send(&calculatorpb.FindMaxResponse{
			Max: max,
		})

		//Handle error while sending response
		if nil != err {
			log.Fatalf("Error while sending max: %v", err)
			return err
		}
	}
}

func (server *Server) Sqrt(ctx context.Context, request *calculatorpb.SqrtRequest) (*calculatorpb.SqrtResponse, error) {

	//Logging
	log.Println("Square Root API called ...")

	//Fetch data from request
	number := request.GetNumber()

	//Handle for case: square root for a negative number
	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Expected number greater or equal 0",
		)
	}

	return &calculatorpb.SqrtResponse{Sqrt: math.Sqrt(float64(number))}, nil
}

func (server *Server) SumWithDeadline(ctx context.Context, request *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {

	//Logging
	log.Println("Sum With Deadline API called ...")

	//Sleep `sleepCount` seconds!
	//Simulate a heavyweight process
	sleepCount := 3
	for i := 0; i < sleepCount; i++ {

		if context.Canceled == ctx.Err() {
			log.Println("Context is cancelled ...")
			return nil, status.Errorf(
				codes.Canceled,
				"Client cancelled request",
			)
		}

		//Sleep 1 second
		time.Sleep(1 * time.Second)
	}

	response := &calculatorpb.SumResponse{
		Result: request.GetNum1() + request.GetNum2(),
	}

	return response, nil
}

func main() {

	//Setup connection
	listener, err := net.Listen("tcp", IP+":"+PORT)
	if nil != err {
		log.Fatalf("Error while create listen %v", err)
	}

	//SSL File parsing
	certFile := "calculator/ssl/server.crt"
	keyFile := "calculator/ssl/server.key"
	credential, sslError := credentials.NewServerTLSFromFile(certFile, keyFile)
	if nil != sslError {
		log.Fatalf("Create crediential ssl error: %v", sslError)
	}
	options := grpc.Creds(credential)

	//Building the server
	server := grpc.NewServer(options)
	calculatorpb.RegisterCalculatorServer(server, &Server{})
	fmt.Println("Calculator is running")
	err = server.Serve(listener)
	if nil != err {
		log.Fatalf("Error while serve %v", err)
	}
}
