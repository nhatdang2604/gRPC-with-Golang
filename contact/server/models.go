package main

import (
	"log"

	"github.com/beego/beego/v2/client/orm"
)

type ContactInfo struct {
	Id          int64  `orm:"auto"`
	PhoneNumber string `orm:"size(15)"`
	Name        string
	Address     string `orm:"type(text)"`
}

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
