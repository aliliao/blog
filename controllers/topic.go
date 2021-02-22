package controllers

import (
	"blog/models"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	isLogin := IsLogin(this.Ctx.Request)
	topics, err := models.FetchAllTopic(false)
	if err == nil {
		this.Data["Topics"] = topics
	}

	// TODO: 删除文章权限管理
	values, _ := this.Input()
	op := values.Get("op")
	if op == "del" {
		id := values.Get("id")
		if len(id) != 0 {
			err = models.DelTopic(id)
			if err != nil {
				logs.Error(err)
			}
			this.Redirect("/topic", 302)
			return
		}
	}

	this.Data["IsTopic"] = true
	this.Data["Title"] = "文章 - 我的博客"
	this.Data["IsLogin"] = isLogin
	this.TplName = "topic.html"
}

func (this *TopicController) Post() {
	if !IsLogin(this.Ctx.Request) {
		this.Redirect("/signup", 302)
		return
	}
	ck, err := this.Ctx.Request.Cookie("uname")
	if err != nil {
		logs.Error(err)
	}
	author := ck.Value

	values, _ := this.Input()
	title := values.Get("title")
	category := values.Get("category")
	content := values.Get("content")
	tid := values.Get("tid")
	isMarkdown := values.Get("markdown") == "true"

	if len(tid) != 0 {
		err = models.UpdateTopic(tid, title, category, content, isMarkdown)
	} else {
		err = models.AddTopic(title, category, content, author, isMarkdown)
	}
	if err != nil {
		logs.Error(err)
	}

	this.Data["IsLogin"] = true
	this.TplName = "topic_add.html"
	this.Redirect("/topic", 302)
}

func (this *TopicController) Add() {
	if !IsLogin(this.Ctx.Request) {
		this.Redirect("/signup", 302)
		return
	}

	this.Data["Title"] = "新增 - 我的博客"
	this.Data["IsLogin"] = true
	this.TplName = "topic_add.html"
}

func (this *TopicController) View() {
	tid := this.Ctx.Input.Param("0")
	topic, err := models.FetchTopicById(tid)
	if err != nil {
		logs.Error(err)
		this.Redirect("/", 404)
		return
	}

	// Markdown supporting
	if topic.IsMarkdown {
		extensions := parser.CommonExtensions | parser.AutoHeadingIDs
		parser := parser.NewWithExtensions(extensions)
		md := []byte(topic.Content)
		output := markdown.ToHTML(md, parser, nil)
		topic.Content = string(output)
	}

	// comments display
	comments, err := models.FetchAllComments(tid)
	if err == nil {
		this.Data["Comments"] = comments
	} else {
		logs.Error(err)
	}

	this.Data["Topic"] = topic
	this.Data["IsLogin"] = IsLogin(this.Ctx.Request)
	this.TplName = "topic_view.html"
}

func (this *TopicController) Modify() {
	values, _ := this.Input()
	tid := values.Get("tid")
	topic, err := models.FetchTopicById(tid)
	if err != nil {
		logs.Error(err)
		this.Redirect("/", 404)
		return
	}

	this.Data["Topic"] = topic
	this.Data["Title"] = "修改 - 我的博客"
	this.Data["IsLogin"] = IsLogin(this.Ctx.Request)
	this.TplName = "topic_modify.html"
}
