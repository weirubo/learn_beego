package routers

import (
	beego "github.com/astaxie/beego/server/web"
	"web/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    beego.Router("/user/get", &controllers.UserController{}, "*:Read")
}
