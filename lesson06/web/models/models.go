package models

import (
	"github.com/astaxie/beego/client/orm"
	"log"
	"time"
)
import _ "github.com/go-sql-driver/mysql"

type User struct {
	Id      uint64    `form:"-" orm:"auto;pk;description(主键 ID)"`
	Name    string    `form:"name" orm:"unique;size(30)";description(用户名)`
	Age     uint8     `orm:"index;description(年龄)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

// 安装 orm
// go get -u github.com/astaxie/beego/client/orm
// 初始化
func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL) // 导入 mysql 驱动 _ "github.com/go-sql-driver/mysql"
	orm.RegisterDataBase("default", "mysql", "root:root@/beego_web?charset=utf8&loc=Local")
	orm.SetMaxIdleConns("default", 10)
	orm.SetMaxOpenConns("default", 10)
	// orm.RegisterModel(new(User))
	orm.RegisterModelWithPrefix("beego_", new(User)) // 注册模型，可同时注册多个

	// 自动建表
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
}

// 自定义表名
func (u *User) TableName() string {
	return "user"
}

// 自定义存储引擎
func (u *User) TableEngine() string {
	return "INNODB"
}
