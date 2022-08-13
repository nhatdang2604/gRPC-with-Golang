package main

import (
	"log"
	"net"
	"strings"

	"github.com/nhatdang2604/gRPC-with-Golang/gateway/gatewaypb"
	"google.golang.org/grpc"
)

type Server struct{}

const (
	IP   = "0.0.0.0"
	PORT = "50080"
)

func main() {
	listener, err := net.Listen("tcp", strings.Join([]string{IP, PORT}, ":"))

	if nil != err {
		log.Fatalf("Error while listening: %v", err)
	}

	server := grpc.NewServer()
	gatewaypb.RegisterDemoGatewayServer(server, &Server{})
	log.Println("Gateway Server is running")
	err = server.Serve(listener)

	if nil != err {
		log.Fatalf("Error while serving: %v", err)
	}
}
