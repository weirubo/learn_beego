package routers

import (
	beego "github.com/astaxie/beego/server/web"
	"web/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    // 自定义路由
    beego.Router("/user/register", &controllers.UserController{})
    beego.Router("/user/add", &controllers.UserController{}, "*:Create")
    beego.Router("/user/edit", &controllers.UserController{}, "*:Update")
    beego.Router("/user/remove", &controllers.UserController{}, "*:Delete")
}
