package routers

import (
	beego "github.com/astaxie/beego/server/web"
	"github.com/astaxie/beego/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["web/controllers:UserController"] = append(beego.GlobalControllerRouter["web/controllers:UserController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/user/list/:id([0-9]+)",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
