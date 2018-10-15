package controllers

import (
	"github.com/astaxie/beego"
)

type ExportController struct {
	beego.Controller
}

func (this *ExportController) Export() {
	this.Data["Test"] = "报表界面"
	this.TplName = "index.tpl"
}
