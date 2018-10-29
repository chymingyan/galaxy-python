package rules

import (
	"encoding/xml"
	"io/ioutil"
)

type OracleCommand struct {
	//uuid
	OracleCmdId string `xml:"oracleCmdId"`
	//命令名称
	CommandName string `xml:"commandName"`
	//命令文本
	CommandText string `xml:"commandText"`
	//集群命令文本 sql必填项
	RacCommandText string `xml:"racCommandText"`
	//命令结果保存到表的名称
	SaveTableName string `xm:"saveTableName"`
	//命令类型（db命令和shell命令）
	CommandType string `xml:"shell"`
	//Oracle版本
	OracleVersion string `xml:"oracleVersion"`
	//是否是集群
	IsRac bool `xml:"isRac"`
	//是否导出报告
	IsReport bool `xml:"isReport"`
	//是否执行
	IsExec bool `xml:"isExec"`
	//主机类型（Aix Linux）shell必填项
	HostType string `xml:"hostType"`
	//主机系统版本 shell必填项
	OsVersion string `xml:"osVersion"`
}

type OracleCommands struct {
	XMLName    xml.Name        `xml:"oracleRules"`
	Version    string          `xml:"version,attr"`
	Encoding   string          `xml:"encoding,attr"`
	OracleRule []OracleCommand `xml:"oracleRule"`
}

//查询所有Oracle Command命令
func (this *OracleCommands) Commands() (rules []OracleCommand) {
	rules = DatabaseRules.OracleRule
	return
}

//查询单条Command信息
func (this *OracleCommand) Command(cmdId string) (cmd OracleCommand) {
	cmd = OracleCommand{}
	//	ruleCount := len(DatabaseRules.OracleRule)
	//	for i := 0; i < ruleCount; i++ {
	//		if DatabaseRules.OracleRule[i].OracleCmdId == cmdId {
	//			rule = tempRuels, DatabaseRules.OracleRule[i]
	//			break
	//		}
	//	}
	return cmd
}

//添加一条Command命令
func (this *OracleCommands) AddCmd(cmd OracleCommand) bool {
	var isOk = true
	DatabaseRules.OracleRule = append(DatabaseRules.OracleRule, cmd)
	output, err := xml.MarshalIndent(DatabaseRules, "  ", "  ")
	if err != nil {
		isOk = false
	}
	err = ioutil.WriteFile(dbRulesFileName, output, 0666) //写入文件(字节数组)
	if err != nil {
		isOk = false
	}
	//		file, err := os.Open(DbRuleFileName)
	//		if err != nil {
	//			isOk = false
	//		}
	//		defer file.Close()

	//	//写入文件(字节数组)
	//	n2, err := file.Write(output)
	//	if err != nil {
	//		isOk = false
	//	}
	//	return isOk
	return true
}

//修改一条Command命令
func (this *OracleCommand) ModifyCmd(cmd OracleCommand) bool {
	//	tempRuels := &OracleCommand{}
	//	ruleCount := len(DatabaseRules.OracleRule)
	//	for i := 0; i < ruleCount; i++ {
	//		if DatabaseRules.OracleRule[i].OracleCmdId != cmd.OracleCmdId {
	//			tempRuels = append(tempRuels, DatabaseRules.OracleRule[i])
	//		}
	//	}
	//	tempRuels = append(tempRuels, cmd)
	//	DatabaseRules.OracleRule = tempRules
	return true
}

//删除Command命令
func (this *OracleCommands) DelCmd(cmdId string) bool {

	//	tempRules := &OracleCommands{}
	//	ruleCount := len(DatabaseRules.OracleRule)
	//	for i := 0; i < ruleCount; i++ {
	//		if DatabaseRules.OracleRule[i].OracleCmdId != cmdId {
	//			//tempRules = append(tempRules, DatabaseRules.OracleRule[i])
	//		}
	//	}
	//	//DatabaseRules = tempRules
	return true
}
