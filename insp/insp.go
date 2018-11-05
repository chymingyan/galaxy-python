package insp

import (
	"fmt"
)

//获取全部的巡检数据库
func AllInspDb() (dbs []OracleDb) {

}

//通过目标数据库ID获取对应的主机集合
func GetInspHostByDbId(dbId string) (host []OracleHost) {

}

//开始巡检
func Inspection(dbIds []string) bool {

}

//Aix主机巡检
func aixHostInsp() bool {

}

//Linux主机巡检
func linuxHostInsp() bool {

}

//oracle数据库巡检
func oraceInsp() bool {

}
