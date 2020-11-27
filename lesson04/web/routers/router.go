package routers

import (
	"web/controllers"
	beego "github.com/astaxie/beego/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    // beego.AutoRouter(&controllers.UserController{})

    ns := beego.NewNamespace("/v1",
    	beego.NSNamespace("/user",
    		beego.NSRouter("/add", &controllers.UserController{}, "*:Create"),
    		),
    	)

    beego.AddNamespace(ns)
}
