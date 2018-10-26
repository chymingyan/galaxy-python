package command

type AixCommand struct {
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

type AixCommands struct {
	XMLName  xml.Name     `xml:"oracleRules"`
	Version  string       `xml:"version,attr"`
	Encoding string       `xml:"encoding,attr"`
	Xsd      string       `xml:"xmlns:xsd,attr"`
	Xsi      string       `xml:"xmlns:xsi,attr"`
	AixRule  []AixCommand `xml:"oracleRule"`
}
