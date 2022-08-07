package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/nhatdang2604/gRPC-with-Golang/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
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
	}

	log.Printf("Sum API responsed: %v", response.GetResult())
}

func callPND(client calculatorpb.CalculatorClient) {
	log.Println("Calling Prime Number Decomposition API")

	stream, err := client.PrimeNumberDecomposition(context.Background(), &calculatorpb.PNDRequest{
		Number: 242148274,
	})

	if nil != err {
		log.Fatalf("Call Prime Number Decomposition API error: %v", err)
	}

	for {
		response, recvErr := stream.Recv()
		if io.EOF == recvErr {
			log.Println("Server finish streaming")
			break
		}

		log.Printf("Prime number %v", response.GetNumber())
	}
}

func callAverage(client calculatorpb.CalculatorClient) {
	log.Println("Calling Average API")

	stream, err := client.Average(context.Background())

	if nil != err {
		log.Fatalf("Error while calling Average API: %v", err)
	}

	requests := []*calculatorpb.AverageRequest{
		&calculatorpb.AverageRequest{Number: 2},
		&calculatorpb.AverageRequest{Number: 4},
		&calculatorpb.AverageRequest{Number: 6},
		&calculatorpb.AverageRequest{Number: 8},
	}

	for _, request := range requests {
		err = stream.Send(request)

		if nil != err {
			log.Fatalf("Error while sending Average request: %v", err)
		}
	}

	response, err := stream.CloseAndRecv()
	if nil != err {
		log.Fatalf("ERror while recieveing Average response: %v", err)
	}

	log.Printf("Average response %v", response.GetResult())

}

func callFindMax(client calculatorpb.CalculatorClient) {
	log.Println("Calling Find Max API...")

	stream, err := client.FindMax(context.Background())

	if nil != err {
		log.Fatalf("Error while calling Find Max API: %v", err)
	}

	waitChannel := make(chan struct{})

	//A go routine for sending multiple requests
	go func() {

		requests := []*calculatorpb.FindMaxRequest{
			&calculatorpb.FindMaxRequest{Number: 1},
			&calculatorpb.FindMaxRequest{Number: -1},
			&calculatorpb.FindMaxRequest{Number: 2},
			&calculatorpb.FindMaxRequest{Number: -2},
			&calculatorpb.FindMaxRequest{Number: 3},
		}

		for _, request := range requests {
			err := stream.Send(request)
			if nil != err {
				log.Fatalf("Send Find Max Request error: %v", err)
				break
			}
		}

		stream.CloseSend()
	}()

	//A go routine for recieving multiple response
	go func() {
		for {
			response, err := stream.Recv()

			if io.EOF == err {
				log.Printf("Ending Find Max API")
				break
			}

			if nil != err {
				log.Fatalf("Recieve Find Max Response error: %v", err)
				break
			}

			log.Printf("Max: %v\n", response.GetMax())
		}

		//Unblock the closured method
		close(waitChannel)
	}()

	//Make the program blocking here,
	//	waiting the unblock signal from the reciever goroutine
	<-waitChannel

}

func callSqrt(client calculatorpb.CalculatorClient, nums ...int32) {

	log.Println("Calling Square Root API")

	for _, num := range nums {
		response, err := client.Sqrt(context.Background(), &calculatorpb.SqrtRequest{
			Number: num,
		})

		if nil != err {
			log.Printf("Call Square Root API error: %v\r\n", err)
			if errorStatus, ok := status.FromError(err); ok {
				log.Printf("Error message: %v", errorStatus.Message())
				log.Printf("Status code: %v", errorStatus.Code())

				if codes.InvalidArgument == errorStatus.Code() {
					log.Printf("InvalidArgument number %v", num)
				}
			}

		} else {
			log.Printf("Square Root API responsed: %v\r\n", response.GetSqrt())
		}
	}
}

func callSumWithDeadline(client calculatorpb.CalculatorClient, timeout time.Duration) {

	log.Println("Calling Sum With Deadline API")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		timeout,
	)

	defer cancel()

	response, err := client.SumWithDeadline(
		ctx,
		&calculatorpb.SumRequest{
			Num1: 2,
			Num2: 3,
		},
	)

	if nil != err {
		statusError, _ := status.FromError(err)
		log.Printf("Error message: %v\r\n", statusError.Err())
		log.Printf("Status code: %v\r\n", statusError.Code())
		return
	}

	log.Printf("Sum With Deadline API response: %v", response.GetResult())

}

func main() {

	certFile := "calculator/ssl/server.crt"
	credential, sslError := credentials.NewClientTLSFromFile(certFile, "")
	if nil != sslError {
		log.Fatalf("Create client credentials ssl error: %v", sslError)
	}

	clientConnection, err := grpc.Dial(IP+":"+PORT, grpc.WithTransportCredentials(credential))
	//clientConnection, err := grpc.Dial(IP+":"+PORT, grpc.WithInsecure())

	//Error handle
	if nil != err {
		log.Fatalf("Error while dial %v", err)
	}

	//Closing the connection after using
	defer clientConnection.Close()

	client := calculatorpb.NewCalculatorClient(clientConnection)

	//callSum(client)
	//callPND(client)
	//callAverage(client)
	//callFindMax(client)
	//callSqrt(client, -2, -1, 0, 1, 2)
	callSumWithDeadline(client, 1*time.Second) //timeout
	callSumWithDeadline(client, 5*time.Second) //not timeout
}
