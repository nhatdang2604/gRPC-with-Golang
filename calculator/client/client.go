package main

import "google.golang.org/grpc"

const (
	IP   = "localhost"
	PORT = "50030"
)

func main() {
	clientConnection, err := grpc.Dial(IP + ":" + PORT)
}
