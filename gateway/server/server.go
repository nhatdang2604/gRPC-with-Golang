package main

import (
	"context"
	"log"
	"net"
	"strings"

	"github.com/nhatdang2604/gRPC-with-Golang/gateway/gatewaypb"
	"google.golang.org/grpc"
)

type Server struct {

	//Embed the unimplemented server
	gatewaypb.UnimplementedDemoGatewayServer
}

const (
	IP   = "0.0.0.0"
	PORT = "50080"
)

func (server *Server) Echo(ctx context.Context, message *gatewaypb.StringMessage) (*gatewaypb.StringMessage, error) {
	log.Println("Echo API is called...")

	//Get message from request
	msg := message.GetMsg()

	//Binding message to the response
	response := &gatewaypb.StringMessage{
		Msg: msg,
	}

	return response, nil
}

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
