package controllers

import (
	"github.com/astaxie/beego/server/web"
)

type UserController struct {
	web.Controller
}

type user struct {
	Id     int64  `form:"-"`
	Name   string `form:"username"`
	Age    uint32 `form:"age"`
	Gender string
}

func (u *UserController) Create() {
	// u.Ctx.WriteString("UserController@Create")
	// 请求数据
	// uid, err := u.GetInt64("uid")
	// if err != nil {
	// 	log.Fatalln(err.Error())
	// 	return
	// }
	// 输出数据
	// uidStr := strconv.FormatInt(uid, 10)
	// u.Ctx.WriteString(uidStr)
	//
	// uname := u.GetString("username")
	// u.Ctx.WriteString(uname)

	uname := u.Input().Get("username")
	u.Ctx.WriteString(uname)

	// form 表单
	// user := user{}
	// err := u.ParseForm(&user)
	// if err != nil {
	// 	log.Fatalln(err.Error())
	// 	return
	// }
	// fmt.Println(user)
	// u.Redirect("/", 200)

	// Request Body
	// user := user{}
	// err := json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	// if err != nil {
	// 	log.Fatalln(err.Error())
	// 	return
	// }

	// 输出
	// fmt.Println(user)
	// u.Data["json"] = map[string]interface{}{"code": 200, "msg": "success"}
	// u.ServeJSON()

	// 数据绑定
	// 基本类型
	// var uid int64
	// err := u.Ctx.Input.Bind(&uid, "uid")
	// if err != nil {
	// 	log.Fatalln(err.Error())
	// 	return
	// }
	// uidStr := strconv.FormatInt(uid, 10)
	// u.Ctx.WriteString(uidStr)

	// 切片
	// language := make([]string, 0, 2)
	// err = u.Ctx.Input.Bind(&language, "language")
	// if err != nil {
	// 	log.Fatalln(err.Error())
	// 	return
	// }
	// fmt.Println(language)

	// 结构体
	// user := user{}
	// err = u.Ctx.Input.Bind(&user, "user")
	// if err != nil {
	// 	log.Fatalln(err.Error())
	// 	return
	// }
	// fmt.Println(user)

	// 上传文件
	// f, h, err := u.GetFile("avatar")
	// if err != nil {
	// 	log.Fatalln(err.Error())
	// 	return
	// }
	// defer f.Close()
	// err = u.SaveToFile("avatar", "static/img/" + h.Filename)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return
	// }
	// u.Redirect("/", 200)

	// 输出数据
	// u.Ctx.WriteString("Hello beego!")
	// u.Data["json"] = map[string]interface{}{"code":200, "msg":"success"}
	// u.ServeJSON()
	// u.Data["jsonp"] = map[string]interface{}{"code":200, "msg":"success"}
	// u.ServeJSONP()
}
