package rules

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

//数据库规则文件名称
const dbRulesFileName = "conf/DatabaseRules.xml"

//数据库规则全局变量数据
var DatabaseRules OracleCommands

type OracleCommand struct {
	//GUID
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
	//备注
	Remark string
}

type OracleCommands struct {
	XMLName    xml.Name        `xml:"oracleRules"`
	Version    string          `xml:"version,attr"`
	Encoding   string          `xml:"encoding,attr"`
	OracleRule []OracleCommand `xml:"oracleRule"`
}

//加载OracleCommands命令
func LoadDataBaseRules() bool {
	fmt.Println("Database file Name: " + dbRulesFileName)
	var isOk = true
	if _, err := os.Stat(dbRulesFileName); os.IsNotExist(err) {
		isOk = false
	}
	if isOk {
		//存在文件读取到全局变量中
		file, err := os.Open(dbRulesFileName)
		if err != nil {
			fmt.Printf("file: %s", err)
			isOk = false
		}
		defer file.Close()
		data, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Printf("data %s", err)
			isOk = false
		}
		err = xml.Unmarshal(data, &DatabaseRules)
		if err != nil {
			isOk = false
		}
		fmt.Printf("DataBaseRules %s", DatabaseRules.Encoding)

	} else {
		//不存在需要创建文件
		file, err := os.Create(dbRulesFileName)
		if err != nil {
			isOk = false
		}
		defer file.Close()
		//初始化空的数据
		DatabaseRules.Version = "1.0"
		DatabaseRules.Encoding = "utf-8"
		output, err := xml.MarshalIndent(DatabaseRules, "  ", "    ")
		err = ioutil.WriteFile(dbRulesFileName, output, 0666) //写入文件(字节数组)
		if err != nil {
			isOk = false
		}

	}
	return isOk
}

//查询所有Oracle Command命令
func (this *OracleCommands) Commands() (rules []OracleCommand) {
	rules = DatabaseRules.OracleRule
	return
}

//查询单条Command信息
func (this *OracleCommands) Command(cmdId string) (cmd OracleCommand) {
	cmd = OracleCommand{}
	ruleCount := len(DatabaseRules.OracleRule)
	for i := 0; i < ruleCount; i++ {
		if DatabaseRules.OracleRule[i].OracleCmdId == cmdId {
			cmd = DatabaseRules.OracleRule[i]
			break
		}
	}
	return
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
	return isOk
}

//修改一条Command命令
func (this *OracleCommands) ModifyCmd(cmd OracleCommand) bool {
	var isOk = true
	tempRules := []OracleCommand{}
	ruleCount := len(DatabaseRules.OracleRule)
	for i := 0; i < ruleCount; i++ {
		if DatabaseRules.OracleRule[i].OracleCmdId != cmd.OracleCmdId {
			tempRules = append(tempRules, DatabaseRules.OracleRule[i])
		}
	}
	tempRules = append(tempRules, cmd)
	DatabaseRules.OracleRule = tempRules
	output, err := xml.MarshalIndent(DatabaseRules, "  ", "  ")
	if err != nil {
		isOk = false
	}
	err = ioutil.WriteFile(dbRulesFileName, output, 0666) //写入文件(字节数组)
	if err != nil {
		isOk = false
	}
	return isOk
}

//删除Command命令
func (this *OracleCommands) DelCmd(cmdId string) bool {
	var isOk = true
	tempRules := []OracleCommand{}
	ruleCount := len(DatabaseRules.OracleRule)
	for i := 0; i < ruleCount; i++ {
		if DatabaseRules.OracleRule[i].OracleCmdId != cmdId {
			tempRules = append(tempRules, DatabaseRules.OracleRule[i])
		}
	}
	DatabaseRules.OracleRule = tempRules
	output, err := xml.MarshalIndent(DatabaseRules, "  ", "  ")
	if err != nil {
		isOk = false
	}
	err = ioutil.WriteFile(dbRulesFileName, output, 0666) //写入文件(字节数组)
	if err != nil {
		isOk = false
	}
	return isOk
}
