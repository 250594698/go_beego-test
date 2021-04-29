package controllers

import (
	_ "fmt"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type baseController struct {
	beego.Controller
}

func (p *baseController) Succes(msg string, url string) {

	p.Ctx.WriteString("<script>alert('"+msg+"');window.location.href ='"+url+"';</script>")
	p.StopRun()
}

func (p *baseController) Error(msg string, url string) {
	p.Ctx.WriteString("<script>alert('"+msg+"');window.history.go(-1);</script>")
	p.StopRun()
}


//获取用户IP地址
func (p *baseController) getClientIp() string {
	s := strings.Split(p.Ctx.Request.RemoteAddr, ":")
	return s[0]
}