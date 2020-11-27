package main

import (
	_ "web/routers"
	beego "github.com/astaxie/beego/server/web"
)

func main() {
	beego.Run()
}

