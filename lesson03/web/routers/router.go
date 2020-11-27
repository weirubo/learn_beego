package routers

import (
	beego "github.com/astaxie/beego/server/web"
	"github.com/astaxie/beego/server/web/context"
	"web/controllers"
)

func init() {
    // 基础路由
    beego.Get("/get", func(ctx *context.Context) {
		ctx.WriteString("get method")
	})

    beego.Post("/post", func(ctx *context.Context) {
		ctx.WriteString("post method")
	})

    beego.Any("/any", func(ctx *context.Context) {
		ctx.WriteString("all method")
	})

	// 固定路由
	beego.Router("/", &controllers.MainController{})
	// beego.Router("/user/list/:id", &controllers.UserController{})

    // 正则路由
    // beego.Router("/user/list/?:id([0-9]+)", &controllers.UserController{})

    // 自动路由
    beego.AutoRouter(&controllers.UserController{})

	// 注解路由
	// beego.BConfig.WebConfig.CommentRouterPath = "web/controllers"
	// beego.Include(&controllers.UserController{})


    // 命名空间
    ns := beego.NewNamespace("/v1",
    	beego.NSRouter("/user/list/:id", &controllers.UserController{}, "*:List"),
    	)
    beego.AddNamespace(ns)
}
