package controllers

import beego "github.com/beego/beego/v2/server/web"

type VideoController struct {
	beego.Controller
}

func (this *VideoController) Get () {
	this.Data["IsVideo"] = true
	this.Data["Title"] = "视频 - 我的博客"
	this.Data["IsLogin"] = IsLogin(this.Ctx.Request)
	this.TplName = "video.html"
}
