package models

import (
	"errors"
	"github.com/beego/beego/v2/client/orm"
)

var (
	ErrAlreadyExists = errors.New("username already exists.")
)


func AddUser(username, pwd string) error {
	u, err := GetUserByUsername(username)
	if err == nil {
		return ErrAlreadyExists
	}

	u.UserName = username
	u.Pwd = pwd

	o := orm.NewOrm()
	_, err = o.Insert(u)
	return err
}

func GetUserByUsername(username string) (*User, error) {
	o := orm.NewOrm()
	user := &User{UserName: username}

	qs := o.QueryTable(user)
	err := qs.Filter("user_name", username).One(user)
	return user, err
}
