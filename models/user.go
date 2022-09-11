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

func CheckIfUserExsist (uid int32) (bool) {
	o := orm.NewOrm()
	
	cnt,err := o.QueryTable("user").Filter("id",uid).Count()
	if err != nil {
		return false
	}
	
	if(cnt > 0){
		return true
	} else{
		return false
	}

}

func init() {
	orm.RegisterModel(new(User))
}