package controllers

import "C"
import (
	"fmt"
	_ "fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"math"
	"shop/models"
	"strconv"
)

type UserController struct {
	baseController
}

type User struct {

	Id int
	Name string  `orm:"size(20)"`
	Mobile string `orm:"size(11)"`
	User_id  int
	Address_desc string  `orm:"size(200)"`

}

type Usertwo struct {

	Id int
	Name string  `orm:"size(20)"`
	Mobile string `orm:"size(11)"`
	User_id  int
	Address_desc string  `orm:"size(200)"`

}

//type Address struct {
//
//	Address_id int
//	User_id  int
//	Address_desc string  `orm:"size(200)"`
//
//}

func init() {
	// 设置数据库基本信息
	_ = orm.RegisterDataBase("default", "mysql", "root:root@tcp(localhost:3306)/java_test?charset=utf8")
	// 第3个参数是是否打印日志（打印sql语句）生成表,第一个参数是别名，第2个参数是否强制更新（除非表结构发生改变），一般设置为false，若为true，则表数据会被清空！
	//_ = orm.RunSyncdb("default", false, true)
}

func (c *UserController) Index() {

	var (
		list []User
		//users []*models.User
		user models.User
	)

	//user.Query().All(&list)

	pagesize:=5   //每页条数
	//pageSetNum:=1
	page_two:=c.Ctx.Input.Param(":page")
	fmt.Println(page_two)
	page,_:=strconv.Atoi(page_two)
	count,_:=user.Query().Count()

	pageCount:=math.Ceil(float64(count)/float64(pagesize))

	//fmt.Println(pageCount,page)

	//data:=list
	//fmt.Println("%#v\n",data)

	qb,_:=orm.NewQueryBuilder("mysql")
	qb.Select("*").From("user").LeftJoin("address").On("user.id=address.user_id").Limit(pagesize).Offset(page-1)
	sql:=qb.String()
	nums,_:=orm.NewOrm().Raw(sql).QueryRows(&list)
	println(nums)
	println(list)

	var pagelist[]int
	for i:=0;i<int(pageCount);i++  {
		pagelist=append(pagelist,i+1)

	}
	c.Data["count"] = count
	c.Data["pageCount"] = pageCount
	c.Data["pageNum"] = page
	c.Data["pagelist"] = pagelist
	c.Data["list_data"] =list
	c.TplName = "user.tpl"
	//json.MarshalIndent(c.Data["list_data"],"","")
}

func (c *UserController)Add()  {

	c.TplName="add.tpl"
	
}

func (c *UserController)Edit()  {

	c.Ctx.Input.Param(":id")
	id_name := c.Ctx.Input.Param(":id")
	i,_:= strconv.Atoi(id_name)
	//fmt.Println(reflect.TypeOf(i))
	user:= models.User{Id:i}
	orm.NewOrm().Read(&user)
	c.Data["list_data"]=user

	//fmt.Println(user.Name)
	c.TplName="edit.tpl"
	
}

func (c *UserController)AddDo()  {

	id,_:=c.GetInt("id")
	user :=models.User{}
	user.Name=c.GetString("name")
	user.Mobile=c.GetString("mobile")
	if id==0 {
		if _,err:=orm.NewOrm().Insert(&user);err != nil{

			//fmt.Println("添加数据失败")
			c.Error("添加数据失败","")
		}else{

			//fmt.Println("添加数据成功")
			c.Succes("添加数据成功","/user/index")

		}
		
	}else{

		user.Id=id
		if _,err:=orm.NewOrm().Update(&user); err!=nil{
			c.Error("编辑数据失败","")

		}else{
			c.Succes("编辑数据成功","/user/index")
		}

	}

	//fmt.Println(c.GetString("name"))

	c.ServeJSONP()
	
}

func (c *UserController)Del(){

	id_name :=c.Ctx.Input.Param(":id")
	i,_:= strconv.Atoi(id_name)

	fmt.Println(i)

	user :=models.User{Id: i}

	if _,err:=orm.NewOrm().Delete(&user);err!=nil{

		c.Error("删除数据失败","")
	}else{
		c.Succes("删除数据成功","/user/index")
	}

	c.ServeJSONP()

}

func (c *UserController) Test() {

	var (
		list []*models.User
		//users []*models.User
		user models.User
	)

	//println(reflect.TypeOf(list))
	//println(reflect.TypeOf(list2))

	//user.Query().All(&list)

	pagesize:=5   //每页条数
	//pageSetNum:=1
	page_two:=c.Ctx.Input.Param(":page")
	fmt.Println(page_two)
	page,_:=strconv.Atoi(page_two)
	count,_:=user.Query().Count()

	pageCount:=math.Ceil(float64(count)/float64(pagesize))

	//fmt.Println(pageCount,page)

	//data:=list
	//fmt.Println("%#v\n",data)

	qb,_:=orm.NewQueryBuilder("mysql")
	qb.Select("*").From("user").LeftJoin("address").On("user.id=address.user_id").Limit(pagesize).Offset(page-1)
	sql:=qb.String()
	//nums,_:=orm.NewOrm().Raw(sql).QueryRows(&list)

	var list2 []Usertwo
	nums,_:=orm.NewOrm().Raw(sql).QueryRows(&list2)
	println(nums)
	//println(list)

	var pagelist[]int
	for i:=0;i<int(pageCount);i++  {
		pagelist=append(pagelist,i+1)

	}
	c.Data["count"] = count
	c.Data["pageCount"] = pageCount
	c.Data["pageNum"] = page
	c.Data["pagelist"] = pagelist
	c.Data["list_data"] =list
	//c.TplName = "user.tpl"
	c.Data["json"] = list2
	c.ServeJSON()
	//json.MarshalIndent(c.Data["list_data"],"","")
}