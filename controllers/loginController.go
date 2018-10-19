package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Index() {
	this.Data["Login"] = "测试登录界面"
	this.TplName = "login.tpl"
}

func (this *LoginController) Login() {
	hostIp := this.GetString("hostIp")
	hostPort := this.GetString("hostPort")
	serviceName := this.GetString("serviceName")
	userName := this.GetString("userName")
	password := this.GetString("password")
	result := struct {
		val string
	}{userName}
	this.Data["json"] = &result
	this.ServeJSON() //响应前端
	beego.Debug("username:", hostIp, hostPort, serviceName, userName, password)
}
