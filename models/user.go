package models

import "github.com/beego/beego/v2/adapter/orm"

type User struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

func init() {
	orm.RegisterModel(new(User))
}