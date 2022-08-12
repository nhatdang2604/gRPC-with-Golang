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

//Callee to Delete a contact to the database
func DeleteContact(client contactpb.ContactServiceClient, id int64) error {

	//Logging
	log.Println("Calling Delete Contact API....")

	//Create the request
	request := &contactpb.DeleteContactRequest{
		Id: id,
	}

	//Send the request
	response, err := client.Delete(
		context.Background(),
		request,
	)

	//Error handling after sending request
	if nil != err {
		log.Printf("Error while calling the Delete Contact API: %v", err)
		return err
	}

	//Print the response
	log.Printf("Delete Contact API responsed: %v", response)

	return nil
}

//Callee to Read a contact from the database
func ReadContact(client contactpb.ContactServiceClient, id int64) error {

	//Logging
	log.Println("Calling Read Contact API....")

	//Create the request
	request := &contactpb.ReadContactRequest{
		Id: id,
	}

	//Send the request
	response, err := client.Read(
		context.Background(),
		request,
	)

	//Error handling after sending request
	if nil != err {
		log.Printf("Error while calling the Read Contact API: %v", err)
		return err
	}

	//Print the response
	log.Printf("Read Contact API responsed: %v", response)

	return nil
}

//Callee to Search a contact from the database
func SearchContact(client contactpb.ContactServiceClient, keyword string) error {

	//Logging
	log.Println("Calling Search Contact API....")

	//Create the request
	request := &contactpb.SearchContactRequest{
		Keyword: keyword,
	}

	//Send the request
	response, err := client.Search(
		context.Background(),
		request,
	)

	//Error handling after sending request
	if nil != err {
		log.Printf("Error while calling the Search Contact API: %v", err)
		return err
	}

	//Print the response
	log.Printf("Search Contact API responsed: %v", response)

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

	//Dummy value for testing Insert/Update API
	// contact := &contactpb.Contact{
	// 	Id:          int64(2),
	// 	PhoneNumber: "111111111111111",
	// 	Address:     "Test",
	// 	Name:        "Test",
	// }

	//Execute Insert API
	// InsertContact(client, contact)

	//Execute Update API
	// contact.Id = 2
	// contact.Name = "Test000"
	// UpdateContact(client, contact)

	//Execute Delete API
	// deletedContactId := int64(2)
	// DeleteContact(client, deletedContactId)

	//Execute Read API
	// readContactId := int64(3)
	// ReadContact(client, readContactId)

	//Execute Search API
	keyword := "Test001"
	SearchContact(client, keyword)
}
