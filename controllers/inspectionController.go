package controllers

import (
	"github.com/astaxie/beego"
)

type InspectionController struct {
	beego.Controller
}

func (this *InspectionController) Insp() {
	this.Data["Website"] = "https://github.com/chymingyan"
	this.Data["Email"] = "chen_haiyan@hotmail.com"
	this.Data["Index"] = "巡检页面"
	this.TplName = "index.tpl"
}
