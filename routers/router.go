package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"myAppNew/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/login",&controllers.LoginController{})
    beego.Router("/category",&controllers.CategoryController{})
    beego.AutoRouter(&controllers.TopicController{})
    beego.Router("/topic",&controllers.TopicController{})
	beego.Router("/reply/add",&controllers.ReplyController{},"post:Add")
	beego.Router("/reply/delete",&controllers.ReplyController{},"get:Delete")
    beego.Router("/attachment/:all",&controllers.AttachController{})
}
