package main

import (
	"context"
	"flag"
	"net/http"
	"strings"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/nhatdang2604/gRPC-with-Golang/gateway/gatewaypb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	IP               = "localhost"
	GRPC_SERVER_PORT = "50080"

	PORT = "8081"
)

var (

	//Command-line options:
	//gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", strings.Join([]string{IP, GRPC_SERVER_PORT}, ":"), "gRPC server endpoint")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	//Register gRPC server endpoint
	//Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	err := gatewaypb.RegisterDemoGatewayHandlerFromEndpoint(
		ctx,
		mux,
		*grpcServerEndpoint,
		options,
	)

	if nil != err {
		return err
	}

	return http.ListenAndServe(":"+PORT, mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); nil != err {
		glog.Fatal(err)
	}
}
