package controllers

import (
	"blog/models"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"strings"
)

type CommentController struct {
	beego.Controller
}

func (this *CommentController) Add() {
	values, _ := this.Input()
	//fmt.Println(values)

	//this.Ctx.WriteString("暂不支持添加评论！comment data:\n")
	//this.Ctx.WriteString(fmt.Sprint(values))
	//this.Ctx.WriteString(this.Ctx.Request.RemoteAddr)

	ip := strings.Split(this.Ctx.Request.RemoteAddr, ":")[0]
	browserInfo := values.Get("browserinfo")
	osInfo := values.Get("osinfo")
	nickname := values.Get("nickname")
	comment := values.Get("comment")
	topicId := values.Get("tid")

	models.AddComment(topicId, nickname, comment, browserInfo, osInfo, ip)
	url := "/topic/view/" + topicId
	this.Redirect(url, 302)
}

func (this *CommentController) Reply() {
	this.Ctx.WriteString("暂不支持回复！敬请期待！！")
}

func (this *CommentController) Del() {
	this.Ctx.WriteString("暂不支持删除评论！敬请期待！！")
}

func (this *CommentController) Thumbsup() {
	// TODO: 不可重复点赞
	values, _ := this.Input()
	tid := values.Get("tid")
	cid := values.Get("cid")
	err := models.ThumbsUp(cid)
	if err != nil {
		logs.Error(err)
	}
	url := "/topic/view/" + tid
	this.Redirect(url, 302)
}