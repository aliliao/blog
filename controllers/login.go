package controllers

import (
	"blog/models"
	"github.com/beego/beego/v2/core/logs"
	"net/http"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

// 1. 注册新用户
type SignUpController struct {
	beego.Controller
}

func (this *SignUpController) Get() {
	values, _ := this.Input()
	msg := values.Get("msg")
	if len(msg) != 0 {
		this.Data["HasSignUpTips"] = true
		this.Data["SignUpTips"] = msg
	}

	this.Data["Title"] = "注册 - 我的博客"
	this.Data["IsSignUp"] = true
	this.TplName = "login.html"
}

func (this *SignUpController) Post() {
	redirectUrl := this.Ctx.Request.URL.String()
	values, err := this.Input()
	if err != nil {
		logs.Error(err)
	}

	uname := strings.TrimSpace(values.Get("uname"))
	uname = strings.ToLower(uname) // 账号名大小写无关
	pwd := values.Get("pwd")
	if len(uname) != 0 && len(pwd) != 0 {
		err = models.AddUser(uname, pwd)
		logs.Info("uname: ", uname, "pwd: ", pwd)
		if err == nil {
			redirectUrl = redirectUrl + "?msg=注册成功！"
		} else if err == models.ErrAlreadyExists {
			redirectUrl = redirectUrl + "?msg=用户名已存在！"
		} else {
			redirectUrl = redirectUrl + "?msg=注册失败，请重试！"
		}
	} else {
		redirectUrl = redirectUrl + "?msg=输入错误！"
	}

	this.Redirect(redirectUrl, 302)
}

// 根据cookie判断用户是否登入
func IsLogin(r *http.Request) bool {
	uname, err := r.Cookie("uname")
	if err != nil {
		return false
	}

	pwd, err := r.Cookie("pwd")
	if err != nil {
		return false
	}

	user, err := models.GetUserByUsername(uname.Value)
	if err != nil {
		return false
	}

	return uname.Value == user.UserName && pwd.Value == user.Pwd
}

// 2. 用户登入
type SignInController struct {
	beego.Controller
}

func (this *SignInController) Get() {
	this.Data["Title"] = "登入 - 我的博客"
	this.Data["IsSignIn"] = true

	// 判断是否是登入失败重定向来的请求
	values, _ := this.Input()
	this.Data["SignInAuthFailed"] = values.Get("auth") == "failed"

	this.TplName = "login.html"
}

func (this *SignInController) Post() {
	values, err := this.Input()
	if err != nil {
		logs.Error(err)
		return
	}
	uname := strings.TrimSpace(values.Get("uname"))
	pwd := values.Get("pwd")
	autoLogin := values.Get("autoLogin") == "true"

	user, err := models.GetUserByUsername(uname)
	if err == nil {
		if uname == user.UserName && pwd == user.Pwd {
			maxAge := 0
			if autoLogin {
				maxAge = 1<<31 - 1
			}
			this.Ctx.SetCookie("uname", uname, maxAge, "/")
			this.Ctx.SetCookie("pwd", pwd, maxAge, "/")
			this.Redirect("/", 302)
		}
	}

	// 登入失败，设置flag，然后重定向到本页
	url := this.Ctx.Request.URL.String() + "?auth=failed"
	this.Redirect(url, 302)
}

// 3. 用户退出
type SignOutController struct {
	beego.Controller
}

func (this *SignOutController) Get() {
	if IsLogin(this.Ctx.Request) {
		this.Ctx.SetCookie("uname", "", -1, "/")
		this.Ctx.SetCookie("pwd", "", -1, "/")
	}
	this.Redirect("/", 302)
}