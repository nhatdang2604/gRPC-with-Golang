package main

import "github.com/beego/beego/v2/client/orm"

func init() {
	orm.RegisterModel(new(ContactInfo))
}

type ContactInfo struct {
	Id          int64  `orm:"auto"`
	PhoneNumber string `orm:"size[15];pk"`
	Name        string
}
