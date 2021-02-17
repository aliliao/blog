package controllers

import beego "github.com/beego/beego/v2/server/web"

type AboutController struct {
	beego.Controller
}

func (this *AboutController) Get () {
	this.Data["Title"] = "关于 - 我的博客"
	this.Data["IsAbout"] = true
	this.Data["IsLogin"] = IsLogin(this.Ctx.Request)
	this.TplName = "about.html"
}
