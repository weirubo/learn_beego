package controllers

import (
	"fmt"
	"github.com/astaxie/beego/client/orm"
	"github.com/astaxie/beego/server/web"
	"log"
	"web/models"
)

type UserController struct {
	web.Controller
}

func (u *UserController) Read() {
	o := orm.NewOrm()

	// 普通查询

	// Read
	// user := &models.User{
	// 	Id: 2,
	// }
	// err := o.Read(user)

	// user := &models.User{
	// 	Name: "Lucy",
	// }
	// err := o.Read(user, "Name")
	// if err != nil {
	// 	log.Fatalln(err.Error())
	// 	return
	// }

	// ReadOrCreate
	// user := &models.User{
	// 	Name: "Alan",
	// 	Age: 37,
	// }
	// created, id, err := o.ReadOrCreate(user,"Name", "Age")
	// if err != nil {
	// 	log.Fatalln(err.Error())
	// 	return
	// }
	// fmt.Printf("Created:%t,id:%d\n", created, id)
	// fmt.Printf("user:%+v\n", user)

	// 高级查询

	// One
	// var user models.User
	// string 类型的表名
	// err := o.QueryTable("beego_user").One(&user)
	// 结构体的地址
	// err := o.QueryTable(&user).One(&user)
	// 使用对象作为表名
	// err := o.QueryTable(new(models.User)).One(&user)
	// 指定返回字段，其他字段返回字段类型的零值
	// err := o.QueryTable(new(models.User)).One(&user, "Id", "Name")
	// if err != nil {
	// 	log.Fatalln(err.Error())
	// 	return
	// }
	// fmt.Printf("user:%+v\n", user)

	// All
	var users []models.User
	// rows, err := o.QueryTable("beego_user").All(&users)
	// 指定返回字段，其他字段返回字段类型的零值
	// rows, err := o.QueryTable("beego_user").All(&users, "Id", "Name")
	// if err != nil {
	// 	log.Fatalln(err.Error())
	// 	return
	// }
	// fmt.Printf("rows:%d users:%+v\n", rows, users)

	// Count 查询记录数
	// num, err := o.QueryTable(new(models.User)).Count()
	// if err != nil {
	// 	log.Fatalln(err.Error())
	// 	return
	// }
	// fmt.Println("num:", num)

	// 条件查询
	// var users []models.User
	// Filter 包含
	// 表达式和操作符
	// 等于
	// err := o.QueryTable(new(models.User)).Filter("id", 2).One(&users)
	// 大于
	// num, err := o.QueryTable(new(models.User)).Filter("id__gt", 9).All(&users)
	// 大于等于
	// num, err := o.QueryTable(new(models.User)).Filter("id__gte", 9).All(&users)
	// 小于
	// num, err := o.QueryTable(new(models.User)).Filter("id__lt", 5).All(&users)
	// 小于等于
	// num, err := o.QueryTable(new(models.User)).Filter("id__lte", 5).All(&users)
	// IN
	// num, err := o.QueryTable(new(models.User)).Filter("id__in", 2, 4).All(&users)
	// isnull (true:isnull / false: is not null)
	// num, err := o.QueryTable(new(models.User)).Filter("id__isnull", false).All(&users)
	// num, err := o.QueryTable(new(models.User)).Filter("id__isnull", true).All(&users)

	// exact 等于（区分字母大小写）
	// num, err := o.QueryTable(new(models.User)).Filter("name__exact", "frank").All(&users)
	// iexact 等于（不区分大小写）
	// num, err := o.QueryTable(new(models.User)).Filter("name__iexact", "frank").All(&users)

	// contains Like（区分大小写）
	// num, err := o.QueryTable(new(models.User)).Filter("name__contains", "frank").All(&users)
	// icontains Like（ 不区分大小写）
	// num, err := o.QueryTable(new(models.User)).Filter("name__icontains", "frank").All(&users)

	// startswith (前置模糊查询，区分大小写)
	// num, err := o.QueryTable(new(models.User)).Filter("name__startswith", "fran").All(&users)
	// istartswith（前置模糊查询，不区分大小写）
	// num, err := o.QueryTable(new(models.User)).Filter("name__istartswith", "fran").All(&users)

	// endswith（后置模糊查询，区分大小写）
	// num, err := o.QueryTable(new(models.User)).Filter("name__endswith", "er").All(&users)
	// iendswith（后置模糊查询，不区分大小写）
	// num, err := o.QueryTable(new(models.User)).Filter("name__iendswith", "er").All(&users)

	// Exclude 排除
	// num, err := o.QueryTable(new(models.User)).Exclude("name__exact", "frank").All(&users)

	// Limit 限制条数
	// num, err := o.QueryTable(new(models.User)).Limit(4).All(&users)

	// Offset 偏移
	// num, err := o.QueryTable(new(models.User)).Offset(4).All(&users)

	// OrderBy 排序 "column" means ASC, "-column" means DESC.
	// num, err := o.QueryTable(new(models.User)).OrderBy("id").All(&users)
	// num, err := o.QueryTable(new(models.User)).OrderBy("-id").All(&users)

	// Distinct 去重
	num, err := o.QueryTable(new(models.User)).Filter("id__gt", 9).Distinct().All(&users, "Age")

	// Exist 是否存在
	// isExisted := o.QueryTable(new(models.User)).Filter("name__exact", "frank1").Exist()
	// fmt.Println("isExisted:", isExisted)

	// if err != nil {
	// 	log.Fatalln(err.Error())
	// 	return
	// }
	// fmt.Printf("user:%+v\n", users)

	// 原生 sql 查询
	// 单条查询
	// var user models.User
	// err := o.Raw("SELECT id,name,age FROM beego_user WHERE id = ?", 2).QueryRow(&user)
	// if err != nil {
	// 	log.Fatalln(err.Error())
	// 	return
	// }
	// fmt.Printf("user:%+v\n", user)
	// 多条查询
	// ids := []int{1,3,5}
	// num, err := o.Raw("SELECT id,name,age FROM beego_user WHERE id IN (?,?,?)", ids).QueryRows(&users)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	fmt.Printf("nums:%d user:%+v\n", num, users)
}