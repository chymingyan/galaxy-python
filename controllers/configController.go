package controllers

import (
	"github.com/astaxie/beego"
)

type ConfigController struct {
	beego.Controller
}

func (this *ConfigController) Config() {
	this.Data["Test"] = "配置界面"
	this.TplName = "index.tpl"
}
