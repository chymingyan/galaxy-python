package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

//登录首页展示
func (this *LoginController) Index() {
	this.Data["Login"] = "测试登录界面"
	this.TplName = "login.tpl"
}

//请求登录
func (this *LoginController) Login() {
	//	获取Post传递的登陆信息
	hostIp := this.GetString("hostIp")
	hostPort := this.GetString("hostPort")
	serviceName := this.GetString("serviceName")
	userName := this.GetString("userName")
	password := this.GetString("password")
	//	写入Seesion
	this.SetSession("HostIP", hostIp)
	this.SetSession("HostPort", hostPort)
	this.SetSession("ServiceName", serviceName)
	this.SetSession("UserName", userName)
	this.SetSession("Password", password)
	result := struct {
		Val string
	}{"success"}
	//	另外一种json格式
	//	this.Data["json"] = map[string]interface{}{"val": 222, "num": 1111}
	this.Data["json"] = &result
	this.ServeJSON() //响应前端
	beego.Debug("username:", &result, hostIp, hostPort, serviceName, password)
}

//请求登出
func (this *LoginController) SignOut() {

}
