package routers

import (
	"blog/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/robots.txt", &controllers.RobotController{})
    beego.Router("/home", &controllers.HomeController{})
    beego.Router("/signup", &controllers.SignUpController{})
	beego.Router("/signin", &controllers.SignInController{})
    beego.Router("/signout", &controllers.SignOutController{})
    beego.Router("/topic", &controllers.TopicController{})
    beego.AutoRouter(&controllers.TopicController{})
    beego.Router("/category", &controllers.CategoryController{})
    beego.Router("/comment", &controllers.CommentController{})
    beego.Router("/comment/add", &controllers.CommentController{}, "post:Add")
    beego.Router("/comment/del", &controllers.CommentController{}, "post:Del")
    beego.Router("/comment/reply", &controllers.CommentController{}, "get:Reply")
    beego.Router("comment/thumbsup", &controllers.CommentController{}, "get:Thumbsup")
    beego.Router("/music", &controllers.MusicController{})
    beego.Router("/video", &controllers.VideoController{})
    beego.Router("/about", &controllers.AboutController{})
}
