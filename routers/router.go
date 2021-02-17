package routers

import (
	"blog/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/home", &controllers.HomeController{})
    beego.Router("/signup", &controllers.SignUpController{})
	beego.Router("/signin", &controllers.SignInController{})
    beego.Router("/signout", &controllers.SignOutController{})
    beego.Router("/topic", &controllers.TopicController{})
    beego.AutoRouter(&controllers.TopicController{})
    beego.Router("/category", &controllers.CategoryController{})
    beego.Router("/music", &controllers.MusicController{})
    beego.Router("/video", &controllers.VideoController{})
    beego.Router("/about", &controllers.AboutController{})
}
