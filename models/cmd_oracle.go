package models

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
