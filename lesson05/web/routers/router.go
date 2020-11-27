package routers

import (
	"web/controllers"
	beego "github.com/astaxie/beego/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
