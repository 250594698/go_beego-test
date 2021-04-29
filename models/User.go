package models

import (
	_ "fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)
type User struct {

	Id int
	Name string  `orm:"size(20)"`
	Mobile string `orm:"size(11)"`

}

type Address struct{
	Address_id int
	User_id *User `orm:"rel(fk)"`
	Address_desc string `orm:"size(200)"`

}

func TableName() string {

	return "user"
}

func init() {
	// 设置数据库基本信息
	//_ = orm.RegisterDataBase("default", "mysql", "root:root@tcp(localhost:3306)/java_test?charset=utf8")
	orm.RegisterModel(new(User))
	// 第3个参数是是否打印日志（打印sql语句）生成表,第一个参数是别名，第2个参数是否强制更新（除非表结构发生改变），一般设置为false，若为true，则表数据会被清空！
	//_ = orm.RunSyncdb("default", false, true)
}

func (m *User) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}