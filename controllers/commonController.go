package controllers

import (
	"github.com/astaxie/beego"
)

type CommonController struct {
	beego.Controller
}

//初始化目标数据库集合（TargetDbs）目标主机集合（TargetHosts）目标分组集合（TargetGroup）
//初始化数据库巡检规则、主机巡检规则
func (this *CommonController) InitLocalDb() {

}

//数据库版本集合（系统内置为10g、11g、12c）
func (this *CommonController) OracleVersionArray() {

}

//Aix系统版本集合（系统内置为Aix 5.*、Aix6.*、Aix7.*）
func (this *CommonController) AixVersionArray() {

}

//Linux系统版本集合（系统内置为Linux 5.*、Linux6.*、Linux7.*）
func (this *CommonController) LinuxVersionArray() {

}
