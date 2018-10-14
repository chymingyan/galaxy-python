package controllers

import (
	"github.com/astaxie/beego"
)

type ExportController struct {
	beego.Controller
}

func (this *ExportController) export() {
	this.TplName = "index.tpl"
}
