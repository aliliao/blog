package controllers

import (
	"blog/models"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	values, err := this.Input()
	if err == nil {
		op := values.Get("op")
		switch op {
		case "add":
			name := values.Get("name")
			if len(name) == 0 {
				break
			}
			logs.Debug("adding category: ", name)
			err = models.AddCategory(name)
			if err != nil {
				logs.Error(err)
			}
			this.Redirect("/category", 302)
			return
		case "del":
			id := values.Get("id")
			if len(id) == 0 {
				break
			}
			err = models.DelCategory(id)
			if err != nil {
				logs.Error(err)
			}
			this.Redirect("/category", 302)
			return
		}
	}

	categories, err := models.FetchAllCategory()
	if err != nil {
		logs.Error(err)
	}
	this.Data["Title"] = "分类 - 我的博客"
	this.Data["Categories"] = categories
	this.Data["IsCategory"] = true
	this.Data["IsLogin"] = IsLogin(this.Ctx.Request)
	this.TplName = "category.html"

	for i, cat := range categories {
		fmt.Println(i, cat)
	}
}
