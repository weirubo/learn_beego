package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/client/orm"
	"github.com/astaxie/beego/server/web"
	"log"
	"web/models"
)

type UserController struct {
	web.Controller
}

func (u *UserController) Get() {
	u.TplName = "user.html"
}

func (u *UserController) Create() {
	web.BConfig.WebConfig.AutoRender = false
	o := orm.NewOrm() // 创建一个 Ormer

	// 写操作
	// 方式 1
	name := u.GetString("name", "unknow")
	age, _ := u.GetUint8("age", 0)
	user := &models.User{
		Name: name,
		Age: age,
	}

	// 方式 2 copyrequestbody = true
	// user := &models.User{}
	// _ = json.Unmarshal(u.Ctx.Input.RequestBody, &user)

	// 方式 3 如果 form 表单名与 struct 字段名不同，需要使用 form 标签
	// user := &models.User{}
	// _ = u.ParseForm(user) // 传入的 struct 必须是指针

	// 方式 4 数据绑定
	// http://localhost:8080/user/add?profile.Name=ruby&profile.Age=16
	// user := &models.User{}
	// _ = u.Ctx.Input.Bind(&user, "profile")

	id, err := o.Insert(user)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	fmt.Println("id:", id)
	// u.Redirect("/user/register", 302)

	// 批量插入 InsertMulti()
	// users := []models.User{
	// 	{
	// 		Name: "joy",
	// 		Age: 19,
	// 	},
	// 	{
	// 		Name: "kitty",
	// 		Age: 20,
	// 	},
	// }
	users := []models.User{}
	_ = json.Unmarshal(u.Ctx.Input.RequestBody, &users)
	rows, err := o.InsertMulti(2, users)
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println("rows:", rows)
}

func (u *UserController) Update() {
	web.BConfig.WebConfig.AutoRender = false
	o := orm.NewOrm()
	user := models.User{}
	_ = json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	rows, err := o.Update(&user, "Name")
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println("rows:", rows)
}

func (u *UserController) Delete() {
	web.BConfig.WebConfig.AutoRender = false
	o := orm.NewOrm()
	id, _ := u.GetUint64("id", 0)
	user := models.User{
		Id: id,
	}
	rows, err := o.Delete(&user)
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println("rows:", rows)
}

