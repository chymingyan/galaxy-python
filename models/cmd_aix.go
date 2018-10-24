package models

type AixHostCommand struct {
	//uuid
	AixCommandId string
	//命令名称
	CommandName string
	//命令文本
	CommandText string
	//主机类型（aix linux）
	HostType string
	//系统版本
	OsVersion string
	//Oracle 版本
	OracleVersion string
	//是否执行
	IsExec bool
	//备注
	Remark string
}
