package controllers

import (
	"github.com/astaxie/beego"
)

type ConfigController struct {
	beego.Controller
}

func (this *ConfigController) config() {
	this.TplName = "index.tpl"
}
