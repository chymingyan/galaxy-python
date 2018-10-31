package target

const OracleDbsFileName = "conf/OracleDatabases.xml"

var OracleDbList OracleDbs

type OracleDb struct {
	DbId       string
	DbIp       string
	DbPort     string
	DbSid      string
	DbUserName string
	DbPassword string
	GroupId    string
	Alias      string
	DbVer      string
	ExpPath    string
	Remark     string
}

type OracleDbs struct {
	XMLName  xml.Name   `xml:"databases"`
	Version  string     `xml:"version,attr"`
	Encoding string     `xml:"encoding,attr"`
	Database []TargetDb `xml:"targetdb"`
}
