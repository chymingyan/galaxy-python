package controllers

import (
	"github.com/astaxie/beego"
)

type ExportController struct {
	beego.Controller
}

func (this *ExportController) Export() {
	this.Layout = "ipp_layout.tpl"
	this.TplName = "main.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Head"] = "ipp_head.tpl"
	this.LayoutSections["Menu"] = "ipp_menu.tpl"
	this.LayoutSections["Scripts"] = "ipp_scripts.tpl"
	this.LayoutSections["Content"] = "ipp_export.tpl"
}
