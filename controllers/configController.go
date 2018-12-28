package controllers

import (
	"github.com/astaxie/beego"
)

type ConfigController struct {
	beego.Controller
}

//初始化Config配置页面
func (this *ConfigController) Config() {
	this.Layout = "ipp_layout.tpl"
	this.TplName = "main.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Head"] = "ipp_head.tpl"
	this.LayoutSections["Menu"] = "ipp_menu.tpl"
	this.LayoutSections["Scripts"] = "ipp_scripts.tpl"
	this.LayoutSections["Content"] = "ipp_config.tpl"
}

//获取全部分组信息
func (this *ConfigController) Groups() {

}

//获取单个分组信息
func (this *ConfigController) Group() {

}

//添加分组信息
func (this *ConfigController) AddGroup() {

}

//修改分组信息
func (this *ConfigController) ModifyGroup() {

}

//删除分组信息
func (this *ConfigController) DelGroup() {

}

//获取全部的目标主机信息
func (this *ConfigController) TargetHosts() {

}

//获取单个目标主机信息
func (this *ConfigController) SingleHost() {

}

//添加目标主机信息
func (this *ConfigController) AddHost() {

}

//修改目标主机信息
func (this *ConfigController) ModifyHost() {

}

//删除目标主机信息
func (this *ConfigController) DelHost() {

}

//测试目标主机信息连接
func (this *ConfigController) TestConnHost() {

}

//获取全部的目标数据库信息
func (this *ConfigController) TargetDbs() {
	this.Data["Website"] = "https://github.com/chymingyan"
	this.Data["Email"] = "chen_haiyan@hotmail.com"
	this.Data["Index"] = "配置页面"
	this.TplName = "index.tpl"
}

//获取单个的目标数据库信息
func (this *ConfigController) SingleDb() {

}

//新增目标数据库信息
func (this *ConfigController) AddDb() {

}

//修改目标数据库信息
func (this *ConfigController) ModifyDb() {

}

//删除目标数据库信息
func (this *ConfigController) DelDb() {

}

//测试目标数据库信息连接
func (this *ConfigController) TestConnDb() {

}
