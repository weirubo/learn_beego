package main

import (
	beego "github.com/astaxie/beego/server/web"
	_ "web/routers"
)

func main() {
	beego.Run()
}

