package models

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type OracleCommand struct {
	//uuid
	OracleCmdId string
	//命令名称
	CommandName string
	//命令文本
	CommandText string
	//集群命令文本 sql必填项
	RacCommandText string
	//命令结果保存到表的名称
	SaveTableName string
	//命令类型（db命令和shell命令）
	CommandType string
	//Oracle版本
	OracleVersion string
	//是否是集群
	IsRac bool
	//是否导出报告
	IsReport bool
	//是否执行
	IsExec bool
	//主机类型（Aix Linux）shell必填项
	HostType string
	//主机系统版本 shell必填项
	OsVersion string
}

type OracleCommands struct {
	XMLName    xml.Name        `xml:"oracleRules"`
	Version    string          `xml:"version,attr"`
	Encoding   string          `xml:"encoding,attr"`
	Xsd        string          `xml:"xmlns:xsd,attr"`
	Xsi        string          `xml:"xmlns:xsi,attr"`
	Server     string          `xml:server`
	OracleRule []OracleCommand `xml:'oracleRule'`
}

//全局变量记录数据库巡检项目
var DatabaseRules OracleCommands
var DbRuleFileName string

//判断文件是否存在  存在返回 true 不存在返回false
func checkFileIsExist(filename string) bool {
	filename = "conf/" + filename
	//给全局变量DbRuleFileName赋值
	DbRuleFileName = fileName
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		//创建文件
		f, err := os.Create(filename)
		if err != nil {
			exist = false
		}
		DatabaseRules = OracleCommands{}
		defer f.Close()
	} else {
		//打开文件
		file, err := os.Open(filename)
		if err != nil {
			exist = false
		}
		defer file.Close()
		//读出[]byte 格式的文件内容
		data, err := ioutil.ReadAll(file)
		if err != nil {
			exist = false
		}
		//声明OracleCommands结构
		DatabaseRules = OracleCommands{}
		//将[]byte格式内容转换到OracleCommands结构中
		err = xml.Unmarshal(data, &DatabaseRules)
		if err != nil {
			exist = false
		}
	}
	return exist

}

//查询所有Oracle Command命令
func (this *OracleCommands) Commands() {

}

//查询单条Command信息
func (this *OracleCommand) Command() {

}

//添加一条Command命令
func (this *OracleCommand) AddCmd(cmd OracleCommand) bool {
	var exist = true
	DatabaseRules.OracleRule = append(DatabaseRules.OracleRule, cmd)
	output, err := xml.MarshalIndent(DatabaseRules, "  ", "  ")
	if err != nil {
		exist = false
	}
	file, err := os.Open(DbRuleFileName)
	if err != nil {
		exist = false
	}
	defer file.Close()
	//写入文件(字节数组)
	n2, err := file.Write(output)
	if err != nil {
		exist = false
	}
	return exist
}

//修改一条Command命令
func (this *OracleCommand) ModifyCmd() {

}

//删除Command命令
func (this *OracleCommand) DelCmd() {

}

func OperationXmlFile(cmd OracleCommands)
