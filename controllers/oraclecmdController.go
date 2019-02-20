package controllers

import (
	"ipp/models/rules"

	"github.com/astaxie/beego"
	tsgutils "github.com/timespacegroup/go-utils"
)

type OracleCmdController struct {
	beego.Controller
}

func (this *OracleCmdController) OraDb() {
	this.Layout = "ipp_layout.tpl"
	this.TplName = "main.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Head"] = "ipp_head.tpl"
	this.LayoutSections["TreeView"] = "ipp_treeview.tpl"
	this.LayoutSections["Scripts"] = "ipp_scripts.tpl"
	this.LayoutSections["Content"] = "ipp_oracle_rule.tpl"
}

//Oracle数据库命令首页
func (this *OracleCmdController) OracleCmdIndex() {
	isLoad := rules.LoadDataBaseRules()
	if isLoad {
		this.Data["json"] = rules.DatabaseRules
		this.ServeJSON() //响应前端
	} else {
		this.Data["json"] = ""
		this.ServeJSON()
	}
}

//获取单条Oracle command 信息
func (this *OracleCmdController) Command() {

}

//获取全部Oracle command 信息
func (this *OracleCmdController) Commands() {

}

//添加一条Oracle command 信息
func (this *OracleCmdController) Add() {
	isAdd := true
	//	获取Post传递的登陆信息
	ruleName := this.GetString("ruleName")
	ruleText := this.GetString("ruleText")
	racRuleText := this.GetString("racRuleText")
	tableName := this.GetString("tableName")
	ruleType := this.GetString("ruleType")
	hostType := this.GetString("hostType")
	oracleVer := this.GetString("oracleVer")
	osVer := this.GetString("osVer")
	isRac, err := this.GetBool("isRac")
	if err != nil {
		this.Data["json"] = "false"
		this.ServeJSON()
	}
	isReport, err := this.GetBool("isReport")
	if err != nil {
		this.Data["json"] = "false"
		this.ServeJSON()
	}
	isExec, err := this.GetBool("isExec")
	if err != nil {
		this.Data["json"] = "false"
		this.ServeJSON()
	}
	//生成UUID
	id := tsgutils.GUID()
	if isAdd {
		dbcmd := &rules.OracleCommand{OracleCmdId: id, CommandName: ruleName,
			CommandText: ruleText, RacCommandText: racRuleText, SaveTableName: tableName,
			CommandType: ruleType, OracleVersion: oracleVer,
			IsRac: isRac, IsReport: isReport, IsExec: isExec,
			HostType: hostType, OsVersion: osVer}
		isAdd = rules.DatabaseRules.AddCmd(*dbcmd)
	}
	if isAdd {
		this.Data["json"] = "true"
	} else {
		this.Data["json"] = "false"
	}
	this.ServeJSON()

}

//修改一条Oracle command 信息
func (this *OracleCmdController) Modify() {

}

//删除一条Oracle command信息
func (this *OracleCmdController) Del() {

}
