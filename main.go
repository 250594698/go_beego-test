package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "shop/routers"
)

func init()  {

	_ = orm.RegisterDataBase("default", "mysql", "root:root@tcp(localhost:3306)/java_test?charset=utf8")
	// 打开调试模式，开发的时候方便查看orm生成什么样子的sql语句
	orm.Debug = true
}

func main() {
	beego.Run()
}

