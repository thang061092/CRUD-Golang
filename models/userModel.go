package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type User struct {
	Id        int    `orm:"null"`
	Email     string `orm:"null"`
	Phone     string `orm:"null"`
	Full_name string `orm:"null"`
}

func init() {
	orm.RegisterModel(new(User))
}

func GetPaginateUser() []*User {
	o := orm.NewOrm()

	var users []*User
	o.QueryTable(new(User)).All(&users)

	return users
}

func GetUserById(id int) *User {
	o := orm.NewOrm()
	u := User{Id: id}
	o.Read(&u)
	return &u
}

func CreateUser(data *User) (int64, error) {
	o := orm.NewOrm()

	id, err := o.Insert(data)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func UpdateUser(id int, data *User) *User {
	o := orm.NewOrm()
	u := User{Id: id}

	if o.Read(&u) == nil {
		u.Full_name = data.Full_name
		o.Update(&u)
	}
	return &u
}
