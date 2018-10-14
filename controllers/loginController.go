package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Login() {
	this.Data["Login"] = "测试登录界面"
	this.TplName = "login.tpl"
}
