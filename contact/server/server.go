package main

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/nhatdang2604/gRPC-with-Golang/contact/contactpb"
	"google.golang.org/grpc"
)

const (
	IP   = "0.0.0.0"
	PORT = "50070"
)

type Server struct{}

func main() {
	listener, err := net.Listen("tcp", strings.Join([]string{IP, PORT}, ":"))

	if nil != err {
		log.Fatalf("Error while creating listener: %v", err)
	}

	server := grpc.NewServer()
	contactpb.RegisterContactServiceServer(server, &Server{})
	fmt.Println("Contact service is running")
	err = server.Serve(listener)

	if nil != err {
		log.Fatalf("Error while serving: %v", err)
	}
}
