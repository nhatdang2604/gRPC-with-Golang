package main

import (
	"log"

	"github.com/beego/beego/v2/client/orm"
	"github.com/nhatdang2604/gRPC-with-Golang/contact/contactpb"
)

type ContactInfo struct {
	Id          int64  `orm:"auto"`
	PhoneNumber string `orm:"size(15)"`
	Name        string
	Address     string `orm:"type(text)"`
}

//Insert the caller contact info to the database
func (info *ContactInfo) Insert() error {
	o := orm.NewOrm()
	id, err := o.Insert(info)
	if nil != err {
		log.Printf("Insert contact error: %v\r\n", err)
		return err
	}

	log.Printf("Insert contact with id = %v successfully", id)
	return nil
}

//Parse the contactpb.Contact to ContactInfo
func Parse(target contactpb.Contact) *ContactInfo {

	//Using target getter to inject to the ContactInfo
	result := ContactInfo{
		Id:          target.GetId(),
		PhoneNumber: target.GetPhoneNumber(),
		Name:        target.GetName(),
		Address:     target.GetAddress(),
	}

	return &result
}
