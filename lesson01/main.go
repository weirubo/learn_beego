package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func main() {
	beego.Get("/", func(ctx *context.Context) {
		ctx.WriteString("Hello beego!")
	})
	beego.Run()
}
