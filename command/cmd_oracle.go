package command

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
	Server     string          `xml:"server"`
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
	ruleCount := len(DatabaseRules.OracleRule)
	for i := 0; j < ruleCount; i++ {
		if DatabaseRules.OracleRule[i].OracleCmdId == cmdId {
			rule = tempRuels, DatabaseRules.OracleRule[i]
			break
		}
	}
	return
}

//添加一条Command命令
func (this *OracleCommand) AddCmd(cmd OracleCommand) bool {
	var isOk = true
	DatabaseRules.OracleRule = append(DatabaseRules.OracleRule, cmd)
	output, err := xml.MarshalIndent(DatabaseRules, "  ", "  ")
	if err != nil {
		isOk = false
	}
	file, err := os.Open(DbRuleFileName)
	if err != nil {
		isOk = false
	}
	defer file.Close()

	//写入文件(字节数组)
	n2, err := file.Write(output)
	if err != nil {
		isOk = false
	}
	return isOk
}

//修改一条Command命令
func (this *OracleCommand) ModifyCmd(cmd OracleCommand) bool {
	tempRuels = &OracleCommand{}
	ruleCount := len(DatabaseRules.OracleRule)
	for i := 0; j < ruleCount; i++ {
		if DatabaseRules.OracleRule[i].OracleCmdId != cmd.OracleCmdId {
			tempRuels = append(tempRuels, DatabaseRules.OracleRule[i])
		}
	}
	tempRuels = append(tempRuels, cmd)
	DatabaseRules.OracleRule = tempRules
}

//删除Command命令
func (this *OracleCommand) DelCmd(cmdId string) bool {
	tempRuels = &OracleCommand{}
	ruleCount := len(DatabaseRules.OracleRule)
	for i := 0; j < ruleCount; i++ {
		if DatabaseRules.OracleRule[i].OracleCmdId != cmdId {
			tempRuels = append(tempRuels, DatabaseRules.OracleRule[i])
		}
	}
	tempRuels = append(tempRuels, cmd)
	DatabaseRules.OracleRule = tempRules
}
