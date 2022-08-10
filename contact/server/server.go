package main

import (
	"context"
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
	STRING_CONNECT_METADATA = "root:dangkl123@tcp(127.0.0.1:3306)/contact?charset=utf8"

	//Errors code from Insert Contact API
	SUCCESS_CODE = iota + 1 //start SUCCESS_CODE with 1
	INSERT_CONTACT_ERROR_CODE
	UPDATE_CONTACT_ERROR_CODE
	DELETE_CONTACT_ERROR_CODE
	DELETE_CONTACT_ERROR_NOT_FOUND_CODE
)

type Server struct{}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	dataSource := STRING_CONNECT_METADATA
	err := orm.RegisterDataBase("default", "mysql", dataSource)
	if nil != err {
		log.Panicf("Connect database failed: %v", err)
	}

	//Register models
	orm.RegisterModel(new(ContactInfo))

	//Automatic generate table
	err = orm.RunSyncdb("default", false, true)
	if nil != err {
		log.Panicf("Migrate database failed: %v", err)
	}

	log.Println("Connect database successfully")
}

//Insert the contact from the request
func (server *Server) Insert(ctx context.Context, request *contactpb.InsertContactRequest) (response *contactpb.InsertContactResponse, err error) {
	log.Println("Insert Contact API is called...")

	contactInfo := Parse(*request.GetContact())
	err = contactInfo.Insert()

	if nil != err {
		response = &contactpb.InsertContactResponse{
			StatusCode: INSERT_CONTACT_ERROR_CODE,
			Message:    "Error while inserting contact",
		}

		return
	}

	response = &contactpb.InsertContactResponse{
		StatusCode: SUCCESS_CODE,
		Message:    "OK",
	}

	return
}

//Update the contact from the request
func (server *Server) Update(ctx context.Context, request *contactpb.UpdateContactRequest) (response *contactpb.UpdateContactResponse, err error) {
	log.Println("Update Contact API is called...")

	contactInfo := Parse(*request.GetContact())
	err = contactInfo.Update()

	if nil != err {
		response = &contactpb.UpdateContactResponse{
			StatusCode: UPDATE_CONTACT_ERROR_CODE,
			Message:    "Error while Updating contact",
		}

		return
	}

	response = &contactpb.UpdateContactResponse{
		StatusCode: SUCCESS_CODE,
		Message:    "OK",
	}

	return
}

//Delete the contact from the request
func (server *Server) Delete(ctx context.Context, request *contactpb.DeleteContactRequest) (response *contactpb.DeleteContactResponse, err error) {
	log.Println("Delete Contact API is called...")

	//Try to find the deleted contact with the given id
	contactInfo, err := Read(request.GetId())

	//Check if the contact was existed
	if nil != contactInfo {
		response = &contactpb.DeleteContactResponse{
			StatusCode: DELETE_CONTACT_ERROR_NOT_FOUND_CODE,
			Message:    "The contact is not existed",
		}

		return
	}

	err = contactInfo.Delete()

	//Error handling after deleting contact
	if nil != err {
		response = &contactpb.DeleteContactResponse{
			StatusCode: DELETE_CONTACT_ERROR_CODE,
			Message:    "Error while Deleting contact",
		}

		return
	}

	response = &contactpb.DeleteContactResponse{
		StatusCode: SUCCESS_CODE,
		Message:    "OK",
	}

	return
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
