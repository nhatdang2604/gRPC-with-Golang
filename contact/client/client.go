package main

import (
	"context"
	"log"
	"strings"

	"github.com/nhatdang2604/gRPC-with-Golang/contact/contactpb"
	"google.golang.org/grpc"
)

const (

	//Network stuffs
	IP   = "localhost"
	PORT = "50070"
)

//Callee to insert a contact to the database
func InsertContact(client contactpb.ContactServiceClient, contact *contactpb.Contact) error {

	//Logging
	log.Println("Calling Insert Contact API....")

	//Create the request
	request := &contactpb.InsertContactRequest{
		Contact: contact,
	}

	//Send the request
	response, err := client.Insert(
		context.Background(),
		request,
	)

	//Error handling after sending request
	if nil != err {
		log.Printf("Error while calling the Insert Contact API: %v", err)
		return err
	}

	//Print the response
	log.Printf("Insert Contact API responsed: %v", response)

	return nil
}

//Callee to Update a contact to the database
func UpdateContact(client contactpb.ContactServiceClient, contact *contactpb.Contact) error {

	//Logging
	log.Println("Calling Update Contact API....")

	//Create the request
	request := &contactpb.UpdateContactRequest{
		Contact: contact,
	}

	//Send the request
	response, err := client.Update(
		context.Background(),
		request,
	)

	//Error handling after sending request
	if nil != err {
		log.Printf("Error while calling the Update Contact API: %v", err)
		return err
	}

	//Print the response
	log.Printf("Update Contact API responsed: %v", response)

	return nil
}

func main() {
	connection, err := grpc.Dial(strings.Join([]string{IP, PORT}, ":"), grpc.WithInsecure())

	if nil != err {
		log.Fatalf("Error while dialing: %v", err)
	}

	//If connect successfully => close after using
	defer connection.Close()

	client := contactpb.NewContactServiceClient(connection)

	//Dummy value for testing API
	contact := &contactpb.Contact{

		PhoneNumber: "111111111111111",
		Address:     "Test",
		Name:        "Test",
	}

	//InsertContact(client, contact)

	contact.Id = 2
	contact.Name = "Test000"
	UpdateContact(client, contact)
}
