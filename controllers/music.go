package controllers

import beego "github.com/beego/beego/v2/server/web"

type MusicController struct {
	beego.Controller
}

func (this *MusicController) Get() {
	this.Data["Title"] = "音乐 - 我的博客"
	this.Data["IsMusic"] = true
	this.Data["IsLogin"] = IsLogin(this.Ctx.Request)
	this.TplName = "music.html"
}
