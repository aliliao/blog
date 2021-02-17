package controllers

import (
	"blog/models"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	topics, err := models.FetchAllTopic(true)
	if err == nil {
		this.Data["Topics"] = topics
	} else {
		logs.Error(err)
	}

	// Markdown supporting
	for i, topic := range topics {
		if topic.IsMarkdown {
			extensions := parser.CommonExtensions | parser.AutoHeadingIDs
			parser := parser.NewWithExtensions(extensions)
			md := []byte(topic.Content)
			output := markdown.ToHTML(md, parser, nil)

			topics[i].Content = string(output)
		}
	}

	this.Data["Title"] = "首页 - 我的博客"
	this.Data["IsHome"] = true
	this.Data["IsLogin"] = IsLogin(this.Ctx.Request)
	this.TplName = "home.html"
}
