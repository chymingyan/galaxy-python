package command

type LinuxCommand struct {
	//uuid
	LinuxCommandId string
	//命令名称
	CommandName string
	//命令文本
	CommandText string
	//主机类型
	HostType string
	//系统版本
	OsVersion string
	//数据库版本
	OracleVersion string
	//是否执行
	IsExec bool
	//备注
	Remark string
}

type LinuxCommands struct {
	XMLName   xml.Name       `xml:"oracleRules"`
	Version   string         `xml:"version,attr"`
	Encoding  string         `xml:"encoding,attr"`
	Xsd       string         `xml:"xmlns:xsd,attr"`
	Xsi       string         `xml:"xmlns:xsi,attr"`
	LinuxRule []LinuxCommand `xml:"oracleRule"`
}
