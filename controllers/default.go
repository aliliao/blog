package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Redirect("/home", 302)

/*
	type user struct {
		Name string
		Age int
		Gender string
	}
	u := user{
		Name: "Ali",
		Age: 18,
		Gender: "Male",
	}
	// 结构体变量
	c.Data["User"] = u

	// slice变量
	nums := []int{1,2,3,4,5,6,7,8}
	c.Data["Nums"] = nums

	// 模板变量
	c.Data["TplVal"] = "Hey What's Up!"

	// html 变量
	c.Data["Html"] = "<h5> Let's Go!Go!Go! </h5>"

	// bool 变量
	c.Data["TrueCond"] = true
	c.Data["FalseCond"] = false

	// pipe
	c.Data["Pipe"] = "<h5> Let's Go!Go!Go! </h5>"

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"

*/

	//appname, _ := config.String("appname")
	//httpport, _ := config.String("httpport")
	//runmode, _ := config.String("runmode")
	//resp := "appname: " + appname +
	//	"\r\nhttpport: " + httpport +
	//	"\r\nrunmode: " + runmode
	//c.Ctx.WriteString(resp)
}
