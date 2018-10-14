package controllers

import (
	"github.com/astaxie/beego"
)

type InspectionController struct {
	beego.Controller
}

func (this *InspectionController) insp() {
	this.TplName = "index.tpl"
}
