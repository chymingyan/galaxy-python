package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	//检查Session
	hostIp := this.GetSession("HostIP")
	if hostIp == nil {
		this.Redirect("/signin", 302)
	}

	this.Data["Test"] = "主界面"
	this.Data["Website"] = "https://github.com/chymingyan"
	this.Data["Email"] = "chen_haiyan@hotmail.com"
	this.TplName = "index.tpl"
}
