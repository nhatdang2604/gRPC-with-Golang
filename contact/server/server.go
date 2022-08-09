package main

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/nhatdang2604/gRPC-with-Golang/contact/contactpb"
	"google.golang.org/grpc"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (

	//Network stuffs
	IP   = "0.0.0.0"
	PORT = "50070"

	//Database stuffs
)

type Server struct{}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	dataSource := "root:dangkl123@tcp(127.0.0.1:3306)/contact?charset=utf8"
	err := orm.RegisterDataBase("default", "mysql", dataSource)
	if nil != err {
		log.Panicf("Connect database failed: %v", err)
	}

	log.Println("Connect database successfully")
}

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
