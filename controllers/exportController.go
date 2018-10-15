package controllers

import (
	"github.com/astaxie/beego"
)

type ExportController struct {
	beego.Controller
}

func (this *ExportController) Export() {
	this.Data["Website"] = "https://github.com/chymingyan"
	this.Data["Email"] = "chen_haiyan@hotmail.com"
	this.Data["Index"] = "报表页面"
	this.TplName = "index.tpl"
}
