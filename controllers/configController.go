package controllers

import (
	"github.com/astaxie/beego"
)

type ConfigController struct {
	beego.Controller
}

func (this *ConfigController) Config() {
	this.Data["Website"] = "https://github.com/chymingyan"
	this.Data["Email"] = "chen_haiyan@hotmail.com"
	this.Data["Index"] = "配置页面"
	this.TplName = "index.tpl"
}
