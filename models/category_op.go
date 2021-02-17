package models

import (
	"github.com/beego/beego/v2/client/orm"
	"strconv"
	"time"
)

func AddCategory(name string) error {
	o := orm.NewOrm()
	cate := &Category{
		Title: name,
		Created: time.Now(),
		TopicTime: time.Now(),
	}
	_, err := o.Insert(cate)
	return err
}

func FetchAllCategory() ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)

	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}

func DelCategory(id string) error {
	o := orm.NewOrm()
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	cate := &Category{Id: cid}
	_, err = o.Delete(cate)
	return err
}
