package controllers

import (
	"github.com/astaxie/beego"
)

type InspectionController struct {
	beego.Controller
}

func (this *InspectionController) Insp() {
	this.Data["Test"] = "巡检界面"
	this.TplName = "index.tpl"
}
