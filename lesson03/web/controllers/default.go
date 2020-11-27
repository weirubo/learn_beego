package controllers

import beego "github.com/astaxie/beego/server/web"

type UserController struct {
	beego.Controller
}

// @router /user/list/:id([0-9]+) [get]
func (u *UserController) List() {
	u.Ctx.WriteString("UserController@List func\n")
	id := u.Ctx.Input.Param(":id")
	u.Ctx.WriteString(id)
	params := u.Ctx.Input.Params()
	u.Ctx.WriteString(params["0"])
	u.Ctx.WriteString(params["1"])
}

func (u *UserController) Create() {
	u.Ctx.WriteString("Create func")
}


type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}