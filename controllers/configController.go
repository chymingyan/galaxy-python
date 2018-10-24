package controllers

import (
	"github.com/astaxie/beego"
)

type ConfigController struct {
	beego.Controller
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
