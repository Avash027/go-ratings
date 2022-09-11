package models

import "github.com/beego/beego/v2/adapter/orm"

type User struct {
	Id   int32  `orm:"pk" json:"id"`
	Name string `json:"name"`
}

func AddUser(user *User) (error) {
	o := orm.NewOrm()
	_,err := o.Insert(user)
	return err;
}

func init() {
	orm.RegisterModel(new(User))
}