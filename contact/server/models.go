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

	log.Printf("Insert contact with id = %v successfully\r\n", id)
	return nil
}

//Update the caller contact info to the database
func (info *ContactInfo) Update() error {
	o := orm.NewOrm()

	//Update all fields of the current info
	_, err := o.Update(info)
	if nil != err {
		log.Printf("Update contact error: %v\r\n", err)
		return err
	}

	log.Printf("Update contact with id = %v successfully\r\n", info.Id)
	return nil
}

//Delete the caller contact from the database
func (info *ContactInfo) Delete() error {
	o := orm.NewOrm()

	_, err := o.Delete(info)
	if nil != err {
		log.Printf("Delete contact error: %v\r\n", err)
		return err
	}

	log.Printf("Delete contact with id = %v successfully\r\n", info.Id)
	return nil
}

//Read the Contact Info with the given id from database
func Read(id int64) (*ContactInfo, error) {
	o := orm.NewOrm()

	info := &ContactInfo{Id: id}
	err := o.Read(info)

	if nil != err {
		log.Printf("Read contact error: %v\r\n", err)
		return nil, err
	}

	log.Printf("Read contact with id = %v successfully\r\n", id)
	return info, nil
}

//Search the Contact Info with the given keyword
//	The keyword would be used to search by name of the contacts' owners's name
func Search(keyword string) ([]*ContactInfo, error) {
	o := orm.NewOrm()

	//Using query setter to query on the table of table object
	tableObject := new(ContactInfo)
	querySetter := o.QueryTable(tableObject)

	//buffer to return result
	infos := []*ContactInfo{}

	//Select * from contact_info where name likes %keyword%
	resultCount, err := querySetter.Filter("name__icontains", keyword).All(&infos)

	//Handle when empty result
	if orm.ErrNoRows == err {
		log.Printf("Not found!\r\n")
		return infos, nil
	}

	//Error handle
	if nil != err {
		log.Printf("Search contact error: %v\r\n", err)
		return nil, err
	}

	//Happy path handle
	log.Printf("Search contact with keyword = %v found %d result(s)\r\n", keyword, resultCount)
	return infos, nil
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
