package routers

import (
	"github.com/astaxie/beego"
	"shop/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/user/index/?:page",&controllers.UserController{},"*:Index")
	beego.Router("/user/test/?:page",&controllers.UserController{},"*:Test")
	beego.Router("/user/add",&controllers.UserController{},"*:Add")
	beego.Router("/user/edit/?:id",&controllers.UserController{},"*:Edit")
	beego.Router("/user/adddo",&controllers.UserController{},"*:AddDo")
	beego.Router("/user/del/?:id",&controllers.UserController{},"*:Del")
}
