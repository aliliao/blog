package controllers

import beego "github.com/beego/beego/v2/server/web"

type RobotController struct {
	beego.Controller
}

func (this *RobotController) Get() {
	this.Ctx.WriteString("robot content")
}
