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

//根据数据库ID获取对应的主机信息
func (this *InspectionController) ShowHostByDbId() {

}
